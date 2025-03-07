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

func TestQuerySimpleWithGroupByNumberWithoutRenderedGroupAndChildCountWithLimit(t *testing.T) {
	test := testUtils.RequestTestCase{
		Description: "Simple query with group by number, no children, count with limit on non-rendered group",
		Request: `query {
					Users(groupBy: [Age]) {
						Age
						_count(_group: {limit: 1})
					}
				}`,
		Docs: map[int][]string{
			0: {
				`{
					"Name": "John",
					"Age": 32
				}`,
				`{
					"Name": "Bob",
					"Age": 32
				}`,
				`{
					"Name": "Alice",
					"Age": 19
				}`,
			},
		},
		Results: []map[string]any{
			{
				"Age":    uint64(32),
				"_count": 1,
			},
			{
				"Age":    uint64(19),
				"_count": 1,
			},
		},
	}

	executeTestCase(t, test)
}

func TestQuerySimpleWithGroupByNumberWithRenderedGroupWithLimitAndChildCountWithLimit(t *testing.T) {
	test := testUtils.RequestTestCase{
		Description: "Simple query with group by number, child limit, count with limit on rendered group",
		Request: `query {
					Users(groupBy: [Age]) {
						Age
						_count(_group: {limit: 1})
						_group (limit: 2) {
							Name
						}
					}
				}`,
		Docs: map[int][]string{
			0: {
				`{
					"Name": "John",
					"Age": 32
				}`,
				`{
					"Name": "Bob",
					"Age": 32
				}`,
				`{
					"Name": "Shahzad",
					"Age": 32
				}`,
				`{
					"Name": "Alice",
					"Age": 19
				}`,
			},
		},
		Results: []map[string]any{
			{
				"Age":    uint64(32),
				"_count": 1,
				"_group": []map[string]any{
					{
						"Name": "Bob",
					},
					{
						"Name": "John",
					},
				},
			},
			{
				"Age":    uint64(19),
				"_count": 1,
				"_group": []map[string]any{
					{
						"Name": "Alice",
					},
				},
			},
		},
	}

	executeTestCase(t, test)
}
