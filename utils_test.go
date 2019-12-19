package mgm_test

import (
	"github.com/stretchr/testify/require"
	"mgm"
	"testing"
)

// Coll return model's collection.
func TestGetModelCollection(t *testing.T) {
	setupDefConnection()

	doc := &Doc{}
	coll := mgm.Coll(doc)
	modelCollection := mgm.CollectionByName(doc.CollectionName())
	require.Equal(t, coll.Name(), modelCollection.Name(), "Expected doc's collection , got %v", )
}
