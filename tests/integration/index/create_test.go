// Copyright 2023 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package index

import (
	"testing"

	"github.com/sourcenetwork/defradb/request/graphql/schema"
	testUtils "github.com/sourcenetwork/defradb/tests/integration"
)

func TestIndexCreateWithCollection_ShouldNotHinderQuerying(t *testing.T) {
	test := testUtils.TestCase{
		Description: "Creation of index with collection should not hinder querying",
		Actions: []any{
			testUtils.SchemaUpdate{
				Schema: `
					type Users {
						Name: String @index
						Age: Int
					}
				`,
			},
			testUtils.CreateDoc{
				CollectionID: 0,
				// bae-52b9170d-b77a-5887-b877-cbdbb99b009f
				Doc: `
					{
						"Name":	"John",
						"Age":	21
					}`,
			},
			testUtils.Request{
				Request: `
					query  {
						Users {
							Name
							Age
						}
					}`,
				Results: []map[string]any{
					{
						"Name": "John",
						"Age":  uint64(21),
					},
				},
			},
		},
	}

	testUtils.ExecuteTestCase(t, test)
}

func TestIndexCreate_ShouldNotHinderQuerying(t *testing.T) {
	test := testUtils.TestCase{
		Description: "Creation of index separately from a collection should not hinder querying",
		Actions: []any{
			testUtils.SchemaUpdate{
				Schema: `
					type Users {
						Name: String 
						Age: Int
					}
				`,
			},
			testUtils.CreateDoc{
				CollectionID: 0,
				// bae-52b9170d-b77a-5887-b877-cbdbb99b009f
				Doc: `
					{
						"Name":	"John",
						"Age":	21
					}`,
			},
			testUtils.CreateIndex{
				CollectionID: 0,
				IndexName:    "some_index",
				FieldName:    "Name",
			},
			testUtils.Request{
				Request: `
					query  {
						Users {
							Name
							Age
						}
					}`,
				Results: []map[string]any{
					{
						"Name": "John",
						"Age":  uint64(21),
					},
				},
			},
		},
	}

	testUtils.ExecuteTestCase(t, test)
}

func TestIndexCreate_IfInvalidIndexName_ReturnError(t *testing.T) {
	test := testUtils.TestCase{
		Description: "If invalid index name is provided, return error",
		Actions: []any{
			testUtils.SchemaUpdate{
				Schema: `
					type Users {
						Name: String 
						Age: Int
					}
				`,
			},
			testUtils.CreateIndex{
				CollectionID:  0,
				IndexName:     "!",
				FieldName:     "Name",
				ExpectedError: schema.NewErrIndexWithInvalidName("!").Error(),
			},
		},
	}

	testUtils.ExecuteTestCase(t, test)
}
