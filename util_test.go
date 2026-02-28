package mgm_test

import (
	"testing"

	"github.com/kamva/mgm/v3"
	"github.com/stretchr/testify/require"
)

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
	require.NotNil(t, option)

	// Verify the builder produces a non-empty options list
	optFuncs := option.List()
	require.NotEmpty(t, optFuncs, "UpsertTrueOption should produce at least one option setter")
}

// CustomCollDoc implements CollectionGetter to return a custom collection.
type CustomCollDoc struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `bson:"name"`
}

func (d *CustomCollDoc) Collection() *mgm.Collection {
	// Return a collection with a custom name
	return mgm.CollectionByName("custom_docs")
}

func TestCollWithCollectionGetter(t *testing.T) {
	setupDefConnection()

	doc := &CustomCollDoc{Name: "test"}
	coll := mgm.Coll(doc)

	require.Equal(t, "custom_docs", coll.Name(),
		"Coll() should use CollectionGetter when implemented")
}

func TestCollNameIgnoresCollectionGetter(t *testing.T) {
	// CollName uses reflection/CollectionNameGetter, not CollectionGetter
	doc := &CustomCollDoc{}
	name := mgm.CollName(doc)
	require.Equal(t, "custom_coll_docs", name,
		"CollName should use reflection, not CollectionGetter")
}

func TestCollWithOptions(t *testing.T) {
	setupDefConnection()

	doc := &Doc{}
	coll := mgm.Coll(doc)

	require.Equal(t, mgm.CollName(doc), coll.Name())
}

func TestCollNamePluralizations(t *testing.T) {
	type Category struct {
		mgm.DefaultModel `bson:",inline"`
	}
	type Person struct {
		mgm.DefaultModel `bson:",inline"`
	}
	type Status struct {
		mgm.DefaultModel `bson:",inline"`
	}

	require.Equal(t, "categories", mgm.CollName(&Category{}))
	require.Equal(t, "people", mgm.CollName(&Person{}))
	require.Equal(t, "statuses", mgm.CollName(&Status{}))
}
