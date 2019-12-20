package mgm_test

import (
	"github.com/Kamva/mgm/internal/util"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestPrepareInvalidId(t *testing.T) {
	d := &Doc{}

	_, err := d.PrepareId("test")
	require.Error(t, err, "Expected get error on invalid id value")
}

func TestPrepareId(t *testing.T) {
	d := &Doc{}

	hexId := "5df7fb2b1fff9ee374b6bd2a"
	val, err := d.PrepareId(hexId)
	id, _ := primitive.ObjectIDFromHex(hexId)
	require.Equal(t, val.(primitive.ObjectID), id)
	util.AssertErrIsNil(t, err)
}
