package mgm_test

import (
	"testing"

	"github.com/kamva/mgm/v3/internal/util"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestPrepareInvalidId(t *testing.T) {
	d := &Doc{}

	_, err := d.PrepareID("test")
	require.Error(t, err, "Expected get error on invalid id value")
}

func TestPrepareId(t *testing.T) {
	d := &Doc{}

	hexId := "5df7fb2b1fff9ee374b6bd2a"
	val, err := d.PrepareID(hexId)
	id, _ := primitive.ObjectIDFromHex(hexId)
	require.Equal(t, val.(primitive.ObjectID), id)
	util.AssertErrIsNil(t, err)
}

func TestVersion(t *testing.T) {
	d := &Doc{}
	require.Equal(t, 0, d.GetVersion())
	require.Equal(t, "_v", d.GetVersionFieldName())
	d.IncrementVersion()
	require.Equal(t, 1, d.GetVersion())
}
