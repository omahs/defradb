// Copyright 2022 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package db

import (
	"context"
	"encoding/json"
	"errors"

	ds "github.com/ipfs/go-datastore"
	dsq "github.com/ipfs/go-datastore/query"
	"github.com/libp2p/go-libp2p/core/peer"

	"github.com/sourcenetwork/defradb/client"
	"github.com/sourcenetwork/defradb/core"
	"github.com/sourcenetwork/defradb/datastore"
)

// setReplicator adds a new replicator to the database.
func (db *db) setReplicator(ctx context.Context, txn datastore.Txn, rep client.Replicator) error {
	existingRep, err := db.getReplicator(ctx, rep.Info)
	if errors.Is(err, ds.ErrNotFound) {
		return db.saveReplicator(ctx, txn, rep)
	}
	if err != nil {
		return err
	}

	newSchemas := []string{}
	for _, newSchema := range rep.Schemas {
		isNew := true
		for _, existingSchema := range existingRep.Schemas {
			if existingSchema == newSchema {
				isNew = false
				break
			}
		}
		if isNew {
			newSchemas = append(newSchemas, newSchema)
		}
	}
	rep.Schemas = append(existingRep.Schemas, newSchemas...)
	return db.saveReplicator(ctx, txn, rep)
}

// deleteReplicator removes a replicator from the database.
func (db *db) deleteReplicator(ctx context.Context, txn datastore.Txn, rep client.Replicator) error {
	if len(rep.Schemas) == 0 {
		return db.deleteReplicatorKey(ctx, txn, rep.Info.ID)
	}
	return db.deleteSchemasForReplicator(ctx, txn, rep)
}

func (db *db) deleteReplicatorKey(ctx context.Context, txn datastore.Txn, pid peer.ID) error {
	key := core.NewReplicatorKey(pid.String())
	return txn.Systemstore().Delete(ctx, key.ToDS())
}

func (db *db) deleteSchemasForReplicator(ctx context.Context, txn datastore.Txn, rep client.Replicator) error {
	existingRep, err := db.getReplicator(ctx, rep.Info)
	if err != nil {
		return err
	}

	updatedSchemaList := []string{}
	for _, s := range existingRep.Schemas {
		found := false
		for _, toDelete := range rep.Schemas {
			if toDelete == s {
				found = true
				break
			}
		}
		if !found {
			updatedSchemaList = append(updatedSchemaList, s)
		}
	}

	if len(updatedSchemaList) == 0 {
		return db.deleteReplicatorKey(ctx, txn, rep.Info.ID)
	}

	existingRep.Schemas = updatedSchemaList
	return db.saveReplicator(ctx, txn, existingRep)
}

// GetAllReplicators returns all replicators of the database.
func (db *db) getAllReplicators(ctx context.Context, txn datastore.Txn) ([]client.Replicator, error) {
	reps := []client.Replicator{}
	// create collection system prefix query
	prefix := core.NewReplicatorKey("")
	results, err := txn.Systemstore().Query(ctx, dsq.Query{
		Prefix: prefix.ToString(),
	})
	if err != nil {
		return nil, err
	}

	for result := range results.Next() {
		var rep client.Replicator
		err = json.Unmarshal(result.Value, &rep)
		if err != nil {
			return nil, err
		}

		reps = append(reps, rep)
	}

	return reps, nil
}

func (db *db) getReplicator(ctx context.Context, info peer.AddrInfo) (client.Replicator, error) {
	rep := client.Replicator{}
	key := core.NewReplicatorKey(info.ID.String())
	value, err := db.systemstore().Get(ctx, key.ToDS())
	if err != nil {
		return rep, err
	}

	err = json.Unmarshal(value, &rep)
	if err != nil {
		return rep, err
	}

	return rep, nil
}

func (db *db) saveReplicator(ctx context.Context, txn datastore.Txn, rep client.Replicator) error {
	key := core.NewReplicatorKey(rep.Info.ID.String())
	repBytes, err := json.Marshal(rep)
	if err != nil {
		return err
	}
	return txn.Systemstore().Put(ctx, key.ToDS(), repBytes)
}
