// Copyright 2022 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package simple

import (
	"testing"

	testUtils "github.com/sourcenetwork/defradb/tests/integration"
)

func TestQuerySimpleWithNumericGreaterThanFilterAndNumericOrderDescending(t *testing.T) {
	test := testUtils.RequestTestCase{
		Description: "Simple query with order & filter",
		Request: `query {
					Users(filter: {Age: {_gt: 30}}, order: {Age: DESC}) {
						Name
						Age
					}
				}`,
		Docs: map[int][]string{
			0: {
				`{
					"Name": "John",
					"Age": 21
				}`,
				`{
					"Name": "Bob",
					"Age": 32
				}`,
				`{
					"Name": "Carlo",
					"Age": 55
				}`,
				`{
					"Name": "Alice",
					"Age": 19
				}`,
			},
		},
		Results: []map[string]any{
			{
				"Name": "Carlo",
				"Age":  uint64(55),
			},
			{
				"Name": "Bob",
				"Age":  uint64(32),
			},
		},
	}

	executeTestCase(t, test)
}
