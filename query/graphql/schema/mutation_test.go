package schema

import (
	"strings"
	"testing"

	gql "github.com/graphql-go/graphql"
	"github.com/stretchr/testify/assert"
)

func TestSimleTypeMutation(t *testing.T) {
	sdl := `
	type Book {
		title: String
		description: String
		rating: Float
	}
	`

	sm, err := NewSchemaManager()
	assert.NoError(t, err)
	_, err = sm.Generator.FromSDL(sdl)
	assert.NoError(t, err)

	mutationType := sm.schema.MutationType()
	assert.NotNil(t, mutationType)
	assert.Len(t, mutationType.Fields(), 4)
	for _, mname := range []string{"create_Book", "update_Book", "delete_Book"} {
		f, ok := mutationType.Fields()[mname]
		assert.True(t, ok)

		if strings.Contains(mname, "create") {
			assert.Equal(t, "Book", f.Type.Name())
			assert.IsType(t, &gql.Object{}, f.Type)
			assert.Len(t, f.Args, 1)
		} else if strings.Contains(mname, "update") {
			assert.Equal(t, "Book", f.Type.Name())
			assert.IsType(t, &gql.List{}, f.Type)
			assert.Len(t, f.Args, 3)
		} else if strings.Contains(mname, "delete") {
			assert.Equal(t, "Book", f.Type.Name())
			assert.IsType(t, &gql.List{}, f.Type)
			assert.Len(t, f.Args, 3)
		} else {
			assert.Fail(t, "Invalid mutation name")
		}
	}
}
