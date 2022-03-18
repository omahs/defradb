// Copyright 2022 Democratized Data Foundation.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package fetcher

import (
	"container/list"
	"context"
	"fmt"

	"github.com/sourcenetwork/defradb/client"
	"github.com/sourcenetwork/defradb/core"
	"github.com/sourcenetwork/defradb/datastore"
	"github.com/sourcenetwork/defradb/db/base"
	"github.com/sourcenetwork/defradb/merkle/crdt"

	"github.com/ipfs/go-cid"
	ds "github.com/ipfs/go-datastore"
	format "github.com/ipfs/go-ipld-format"
	dag "github.com/ipfs/go-merkledag"
	"github.com/pkg/errors"
)

var (
	// interface check
	_ Fetcher = (*VersionedFetcher)(nil)
)

// HistoryFetcher is like the normal DocumentFetcher, except it is able to traverse
// to a specific version in the documents history graph, and return the fetched
// state at that point exactly.
//
// Given the following Document state graph
//
// {} --> V1 --> V2 --> V3 --> V4
//		  ^					   ^
//		  |					   |
// 	Target Version		 Current State
//
//
// A regular DocumentFetcher fetches and returns the state at V4, but the
// VersionsedFetcher would step backwards through the update graph, recompose
// the state at the "Target Version" V1, and return the state at that point.
//
// This is achieved by reconstructing the target state using the given MerkleCRDT
// DAG. Given the Target Version CID, we collect all the individual delta nodes
// in the MerkleDAG, until we reach the initial (genesis) state.
//
// Transient/Ephemeral datastores are intanciated for the lifetime of the
// traversal query, on a per object basis. This should be a basic map based
// ds.Datastore, abstracted into a DSReaderWriter.
//
// The goal of the VersionedFetcher is to implement the same external API/Interface as
// the DocumentFetcher, and to have it return the encoded/decoded document as
// defined in the version, so that it can be used as a drop in replacement within
// the scanNode query planner system.
//
// Current limitations:
// - We can only return a single record from an VersionedFetcher
// 	 instance.
// - We can't query into related sub objects (at the moment, as related objects
//   ids aren't in the state graphs.
// - Probably more...
//
// Future optimizations:
// - Incremental checkpoint/snapshotting
// - Reverse traversal (starting from the current state, and working backwards)
// - Create a effecient memory store for in-order traversal (BTree, etc)
//
// Note: Should we transition this state traversal into the CRDT objects themselves, and not
// within a new fetcher?
type VersionedFetcher struct {
	// embed the regular doc fetcher
	*DocumentFetcher

	txn datastore.Txn
	ctx context.Context

	// Transient version store
	root  ds.Datastore
	store datastore.Txn

	key     core.DataStoreKey
	version cid.Cid

	queuedCids *list.List

	col *client.CollectionDescription
	// @todo index  *client.IndexDescription
	mCRDTs map[uint32]crdt.MerkleCRDT
}

// Init

// Start

func (vf *VersionedFetcher) Init(col *client.CollectionDescription, index *client.IndexDescription, fields []*client.FieldDescription, reverse bool) error {
	vf.col = col
	vf.queuedCids = list.New()
	vf.mCRDTs = make(map[uint32]crdt.MerkleCRDT)

	// run the DF init, VersionedFetchers only supports the Primary (0) index
	vf.DocumentFetcher = new(DocumentFetcher)
	return vf.DocumentFetcher.Init(col, &col.Indexes[0], fields, reverse)

}

// Start serializes the correct state accoriding to the Key and CID
func (vf *VersionedFetcher) Start(ctx context.Context, txn datastore.Txn, spans core.Spans) error {
	if vf.col == nil {
		return errors.New("VersionedFetcher cannot be started without a CollectionDescription")
	}

	if len(spans) != 1 {
		return errors.New("spans must contain only a single entry")
	}

	// For the VersionedFetcher, the spans needs to be in the format
	// Span{Start: DocKey, End: CID}
	dk := spans[0].Start()
	cidRaw := spans[0].End()
	if dk.DocKey == "" {
		return errors.New("spans missing start DocKey")
	} else if cidRaw.DocKey == "" { // todo: dont abuse DataStoreKey/Span like this!
		return errors.New("span missing end CID")
	}

	// decode cidRaw from core.Key to cid.Cid
	// need to remove '/' prefix from the core.Key

	c, err := cid.Decode(cidRaw.DocKey)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Failed to decode CID for VersionedFetcher: %s", cidRaw.DocKey))
	}

	vf.txn = txn
	vf.ctx = ctx
	vf.key = dk
	vf.version = c

	// create store
	root := ds.NewMapDatastore()
	vf.root = root
	vf.store, err = datastore.NewTxnFrom(ctx, root, false) // were going to discard and nuke this later
	if err != nil {
		return err
	}

	if err := vf.seekTo(vf.version); err != nil {
		return fmt.Errorf("Failed seeking state to %v: %w", c, err)
	}

	return vf.DocumentFetcher.Start(ctx, vf.store, nil)
}

func (vf *VersionedFetcher) Rootstore() ds.Datastore {
	return vf.root
}

// Start a fetcher with the needed info (cid embedded in a span)

/*
1. Init with DocKey (VersionedFetched is scoped to a single doc)
2. - Create transient stores (head, data, block)
3. Start with a given Txn and CID span set (length 1 for now)
4. call traverse with the target cid
5.

err := VersionFetcher.Start(txn, spans) {
	vf.traverse(cid)
}
*/

// SeekTo exposes the private seekTo
func (vf *VersionedFetcher) SeekTo(ctx context.Context, c cid.Cid) error {
	err := vf.seekTo(c)
	if err != nil {
		return err
	}

	return vf.DocumentFetcher.Start(ctx, vf.store, nil)
}

// seekTo seeks to the given CID version by steping through the CRDT
// state graph from the beginning to the target state, creating the
// serialized state at the given version. It starts by seeking to the
// closest existing state snapshot in the transient Versioned stores,
// which on the first run is 0. It seeks by iteratively jumping through
// the state graph via the `_head` link.
func (vf *VersionedFetcher) seekTo(c cid.Cid) error {
	// reinit the queued cids list
	vf.queuedCids = list.New()

	// recursive step through the graph
	err := vf.seekNext(c, true)
	if err != nil {
		return err
	}

	// after seekNext is completed, we have a populated
	// queuedCIDs list, and all the necessary
	// blocks in our local store
	// If we are using a batch store, then we need to commit
	if vf.store.IsBatch() {
		if err := vf.store.Commit(vf.ctx); err != nil {
			return err
		}
	}

	// if we have a queuedCIDs length of 0, means we don't need
	// to do any more state serialization

	// for cid in CIDs {
	///
	/// vf.merge(cid)
	/// // Note: we need to determine what state we are "Merging"
	/// // into. This isn't necessary for the base case where we only
	/// // are concerned with generating the Versioned state for a single
	/// // CID, but for multiple CIDs, or if we reuse the transient store
	/// // as a cache, we need to swap out states to the parent of the current
	/// // CID.
	// }
	for ccv := vf.queuedCids.Front(); ccv != nil; ccv = ccv.Next() {
		cc, ok := ccv.Value.(cid.Cid)
		if !ok {
			return errors.New("queueudCids contains an invalid CID value")
		}
		err := vf.merge(cc)
		if err != nil {
			return fmt.Errorf("Failed merging state: %w", err)
		}
	}

	// If we are using a batch store, then we need to commit
	if vf.store.IsBatch() {
		if err := vf.store.Commit(vf.ctx); err != nil {
			return err
		}
	}

	// we now have all the the required state stored
	// in our transient local Version_Index, we now need to
	// transfer it to the Primary_Index.

	// Once all values are transferred, exit with no errors
	// Any future operation can resume using the current PrimaryIndex
	// which is actually the serialized state of the CRDT graph at
	// the exact version

	return nil
}

// seekNext is the recursive iteration step of seekTo, its goal is
// to build the queuedCids list, and to transfer the required
// blocks from the global to the local store.
func (vf *VersionedFetcher) seekNext(c cid.Cid, topParent bool) error {
	// check if cid block exists in the global store, handle err

	// @todo: Find an effecient way to determine if a CID is a member of a
	// DocKey State graph
	// @body: We could possibly append the DocKey to the CID either as a
	// child key, or an instance on the CID key.

	hasLocalBlock, err := vf.store.DAGstore().Has(vf.ctx, c)
	if err != nil {
		return fmt.Errorf("(version fetcher) failed to find block in blockstore: %w", err)
	}
	// skip if we already have it locally
	if hasLocalBlock {
		return nil
	}

	blk, err := vf.txn.DAGstore().Get(vf.ctx, c)
	if err != nil {
		return fmt.Errorf("(version fetcher) failed to get block in blockstore: %w", err)
	}

	// store the block in the local (transient store)
	if err := vf.store.DAGstore().Put(vf.ctx, blk); err != nil {
		return fmt.Errorf("(version fetcher) failed to write block to blockstore : %w", err)
	}

	// add the CID to the queuedCIDs list
	if topParent {
		vf.queuedCids.PushFront(c)
	}

	// decode the block
	nd, err := dag.DecodeProtobuf(blk.RawData())
	if err != nil {
		return fmt.Errorf("(version fetcher) failed to decode protobuf: %w", err)
	}

	// subDAGLinks := make([]cid.Cid, 0) // @todo: set slice size
	l, err := nd.GetNodeLink(core.HEAD)
	// ErrLinkNotFound is fine, it just means we have no more head links
	if err != nil && err != dag.ErrLinkNotFound {
		return fmt.Errorf("(version fetcher) failed to get node link from DAG: %w", err)
	}

	// only seekNext on parent if we have a HEAD link
	if err != dag.ErrLinkNotFound {
		err := vf.seekNext(l.Cid, true)
		if err != nil {
			return err
		}
	}

	// loop over links and ignore head links
	for _, l := range nd.Links() {
		if l.Name == core.HEAD {
			continue
		}

		err := vf.seekNext(l.Cid, false)
		if err != nil {
			return err
		}
	}

	return nil
}

// merge in the state of the IPLD Block identified by CID c into the
// VersionedFetcher state.
// Requires the CID to already exists in the DAGStore.
// This function only works for merging Composite MerkleCRDT objects.
//
// First it checks for the existence of the block,
// then extracts the delta object and priority from the block
// gets the existing MerkleClock instance, or creates one.
//
// Currently we assume the CID is a CompositeDAG CRDT node
func (vf *VersionedFetcher) merge(c cid.Cid) error {
	// get node
	nd, err := vf.getDAGNode(c)
	if err != nil {
		return err
	}

	// first arg 0 is the index for the composite DAG in the mCRDTs cache
	if err := vf.processNode(0, nd, client.COMPOSITE, ""); err != nil {
		return err
	}

	// handle subgraphs
	// loop over links and ignore head links
	for _, l := range nd.Links() {
		if l.Name == core.HEAD {
			continue
		}

		// get node
		subNd, err := vf.getDAGNode(l.Cid)
		if err != nil {
			return err
		}

		fieldID := vf.col.GetFieldKey(l.Name)
		if fieldID == uint32(0) {
			return fmt.Errorf("Invalid sub graph field name: %s", l.Name)
		}
		// @todo: Right now we ONLY handle LWW_REGISTER, need to swith on this and get CType from descriptions
		if err := vf.processNode(fieldID, subNd, client.LWW_REGISTER, l.Name); err != nil {
			return err
		}
	}

	return nil
}

func (vf *VersionedFetcher) processNode(crdtIndex uint32, nd format.Node, ctype client.CType, fieldName string) (err error) {
	// handle CompositeDAG
	mcrdt, exists := vf.mCRDTs[crdtIndex]
	if !exists {
		key, err := base.MakePrimaryIndexKeyForCRDT(*vf.col, ctype, vf.key, fieldName)
		if err != nil {
			return err
		}
		mcrdt, err = crdt.DefaultFactory.InstanceWithStores(vf.store, "", nil, ctype, key)
		if err != nil {
			return err
		}
		vf.mCRDTs[crdtIndex] = mcrdt
		// compositeClock = compMCRDT
	}

	delta, err := mcrdt.DeltaDecode(nd)
	if err != nil {
		return err
	}

	height := delta.GetPriority()
	_, err = mcrdt.Clock().ProcessNode(vf.ctx, nil, nd.Cid(), height, delta, nd)
	return err
}

func (vf *VersionedFetcher) getDAGNode(c cid.Cid) (*dag.ProtoNode, error) {
	// get Block
	blk, err := vf.store.DAGstore().Get(vf.ctx, c)
	if err != nil {
		return nil, fmt.Errorf("Failed to get DAG Node: %w", err)
	}

	// get node
	// decode the block
	return dag.DecodeProtobuf(blk.RawData())
}

func (vf *VersionedFetcher) Close() error {
	vf.store.Discard(vf.ctx)
	if err := vf.root.Close(); err != nil {
		return err
	}

	return vf.DocumentFetcher.Close()
}

func NewVersionedSpan(dockey core.DataStoreKey, version cid.Cid) core.Spans {
	// Todo: Dont abuse DataStoreKey for version cid!
	return core.Spans{core.NewSpan(dockey, core.DataStoreKey{DocKey: version.String()})}
}
