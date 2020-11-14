package mgm_test

import (
	"github.com/kamva/mgm/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

// Coll return model's collection.
func TestGetModelCollection(t *testing.T) {
	setupDefConnection()

	doc := &Doc{}
	coll := mgm.Coll(doc)
	name := mgm.CollName(doc)
	require.Equal(t, coll.Name(), name, "Expected doc's collection , got %v")
}

func TestGetDefaultCollName(t *testing.T) {
	type Book struct {
		mgm.DefaultModel `bson:",inline"`
	}

	type BlogPost struct {
		mgm.DefaultModel `bson:",inline"`
	}

	require.Equal(t, "books", mgm.CollName(&Book{}))
	require.Equal(t, "blog_posts", mgm.CollName(&BlogPost{}))
}

type User struct {
	mgm.DefaultModel `bson:",inline"`
}

func (user *User) CollectionName() string {
	return "my_users"
}

func TestGetSpecifiedCollName(t *testing.T) {
	require.Equal(t, "my_users", mgm.CollName(&User{}))
}

func TestUpsertTrueOption(t *testing.T) {
	option := mgm.UpsertTrueOption()
	upsert := true
	assert.Equal(t, option.Upsert, &upsert)
	assert.Nil(t, option.ArrayFilters)
	assert.Nil(t, option.BypassDocumentValidation)
	assert.Nil(t, option.Collation)
}
