// Copyright 2022 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package test_explain_execute

import (
	"testing"

	testUtils "github.com/sourcenetwork/defradb/tests/integration"
	explainUtils "github.com/sourcenetwork/defradb/tests/integration/explain"
)

func TestExecuteExplainMutationRequestWithDeleteUsingID(t *testing.T) {
	test := testUtils.TestCase{

		Description: "Explain (execute) mutation request with deletion using id.",

		Actions: []any{
			explainUtils.SchemaForExplainTests,

			// Addresses
			create2AddressDocuments(),

			testUtils.ExplainRequest{
				Request: `mutation @explain(type: execute) {
					delete_ContactAddress(ids: ["bae-f01bf83f-1507-5fb5-a6a3-09ecffa3c692"]) {
						city
					}
				}`,

				ExpectedFullGraph: []dataMap{
					{
						"explain": dataMap{
							"executionSuccess": true,
							"sizeOfResult":     1,
							"planExecutions":   uint64(2),
							"deleteNode": dataMap{
								"iterations": uint64(2),
								"selectTopNode": dataMap{
									"selectNode": dataMap{
										"iterations":    uint64(2),
										"filterMatches": uint64(1),
										"scanNode": dataMap{
											"iterations":   uint64(2),
											"docFetches":   uint64(1),
											"fieldFetches": uint64(1),
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	explainUtils.ExecuteTestCase(t, test)
}

func TestExecuteExplainMutationRequestWithDeleteUsingFilter(t *testing.T) {
	test := testUtils.TestCase{

		Description: "Explain (execute) mutation request with deletion using filter.",

		Actions: []any{
			explainUtils.SchemaForExplainTests,

			// Author
			create2AuthorDocuments(),

			testUtils.ExplainRequest{
				Request: `mutation @explain(type: execute) {
					delete_Author(filter: {name: {_like: "%Funke%"}}) {
						name
					}
				}`,

				ExpectedFullGraph: []dataMap{
					{
						"explain": dataMap{
							"executionSuccess": true,
							"sizeOfResult":     1,
							"planExecutions":   uint64(2),
							"deleteNode": dataMap{
								"iterations": uint64(2),
								"selectTopNode": dataMap{
									"selectNode": dataMap{
										"iterations":    uint64(2),
										"filterMatches": uint64(1),
										"scanNode": dataMap{
											"iterations":   uint64(2),
											"docFetches":   uint64(2),
											"fieldFetches": uint64(2),
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	explainUtils.ExecuteTestCase(t, test)
}
