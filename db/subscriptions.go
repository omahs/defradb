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

	"github.com/sourcenetwork/defradb/client"
	"github.com/sourcenetwork/defradb/client/request"
	"github.com/sourcenetwork/defradb/events"
	"github.com/sourcenetwork/defradb/planner"
)

func (db *db) checkForClientSubscriptions(r *request.Request) (
	*events.Publisher[events.Update],
	*request.ObjectSubscription,
	error,
) {
	if len(r.Subscription) == 0 || len(r.Subscription[0].Selections) == 0 {
		// This is not a subscription request and we have nothing to do here
		return nil, nil, nil
	}

	if !db.events.Updates.HasValue() {
		return nil, nil, ErrSubscriptionsNotAllowed
	}

	s := r.Subscription[0].Selections[0]
	if subRequest, ok := s.(*request.ObjectSubscription); ok {
		pub, err := events.NewPublisher(db.events.Updates.Value(), 5)
		if err != nil {
			return nil, nil, err
		}

		return pub, subRequest, nil
	}

	return nil, nil, client.NewErrUnexpectedType[request.ObjectSubscription]("SubscriptionSelection", s)
}

func (db *db) handleSubscription(
	ctx context.Context,
	pub *events.Publisher[events.Update],
	r *request.ObjectSubscription,
) {
	for evt := range pub.Event() {
		txn, err := db.NewTxn(ctx, false)
		if err != nil {
			log.Error(ctx, err.Error())
			continue
		}
		defer txn.Discard(ctx)

		p := planner.New(ctx, db.WithTxn(txn), txn)

		s := r.ToSelect(evt.DocKey, evt.Cid.String())

		result, err := p.RunSubscriptionRequest(ctx, s)
		if err != nil {
			pub.Publish(client.GQLResult{
				Errors: []error{err},
			})
			continue
		}

		// Don't send anything back to the client if the request yields an empty dataset.
		if len(result) == 0 {
			continue
		}

		pub.Publish(client.GQLResult{
			Data: result,
		})
	}
}
