package mgm_test

import (
	"github.com/stretchr/testify/require"
	"mgm"
	"testing"
)

// ModelCollection return model's collection.
func TestGetModelCollection(t *testing.T) {
	setupDefConnection()

	model := &Doc{}
	col := model.Collection()
	modelCol := mgm.ModelCollection(model)
	require.Equal(t, col.Name(), modelCol.Name(), "Expected model's collection , got %v", )
}
