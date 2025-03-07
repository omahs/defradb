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

func TestQuerySimple(t *testing.T) {
	test := testUtils.RequestTestCase{
		Description: "Simple query with no filter",
		Request: `query {
					Users {
						_key
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
			},
		},
		Results: []map[string]any{
			{
				"_key": "bae-52b9170d-b77a-5887-b877-cbdbb99b009f",
				"Name": "John",
				"Age":  uint64(21),
			},
		},
	}

	executeTestCase(t, test)
}

func TestQuerySimpleWithAlias(t *testing.T) {
	test := testUtils.RequestTestCase{
		Description: "Simple query with alias, no filter",
		Request: `query {
					Users {
						username: Name
						age: Age
					}
				}`,
		Docs: map[int][]string{
			0: {
				`{
					"Name": "John",
					"Age": 21
				}`,
			},
		},
		Results: []map[string]any{
			{
				"username": "John",
				"age":      uint64(21),
			},
		},
	}

	executeTestCase(t, test)
}

func TestQuerySimpleWithMultipleRows(t *testing.T) {
	test := testUtils.RequestTestCase{
		Description: "Simple query with no filter, multiple rows",
		Request: `query {
					Users {
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
				"Age": 27
			}`,
			},
		},
		Results: []map[string]any{
			{
				"Name": "Bob",
				"Age":  uint64(27),
			},
			{
				"Name": "John",
				"Age":  uint64(21),
			},
		},
	}

	executeTestCase(t, test)
}

func TestQuerySimpleWithUndefinedField(t *testing.T) {
	test := testUtils.RequestTestCase{
		Description: "Simple query for undefined field",
		Request: `query {
					Users {
						Name
						ThisFieldDoesNotExists
					}
				}`,
		ExpectedError: "Cannot query field \"ThisFieldDoesNotExists\" on type \"Users\".",
	}

	executeTestCase(t, test)
}

func TestQuerySimpleWithSomeDefaultValues(t *testing.T) {
	test := testUtils.RequestTestCase{
		Description: "Simple query with some default-value fields",
		Request: `query {
					Users {
						Name
						Email
						Age
						HeightM
						Verified
					}
				}`,
		Docs: map[int][]string{
			0: {
				`{
					"Name": "John"
				}`,
			},
		},
		Results: []map[string]any{
			{
				"Name":     "John",
				"Email":    nil,
				"Age":      nil,
				"HeightM":  nil,
				"Verified": nil,
			},
		},
	}

	executeTestCase(t, test)
}

func TestQuerySimpleWithDefaultValue(t *testing.T) {
	test := testUtils.RequestTestCase{
		Description: "Simple query with default-value fields",
		Request: `query {
					Users {
						Name
						Email
						Age
						HeightM
						Verified
					}
				}`,
		Docs: map[int][]string{
			0: {
				`{ }`,
			},
		},
		Results: []map[string]any{
			{
				"Name":     nil,
				"Email":    nil,
				"Age":      nil,
				"HeightM":  nil,
				"Verified": nil,
			},
		},
	}

	executeTestCase(t, test)
}
