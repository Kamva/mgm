package mgm

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestDateFieldsCreating(t *testing.T) {
	df := &DateFields{}

	before := time.Now().UTC()
	err := df.Creating()
	after := time.Now().UTC()

	require.NoError(t, err)
	assert.False(t, df.CreatedAt.IsZero(), "CreatedAt should be set")
	assert.True(t, !df.CreatedAt.Before(before), "CreatedAt should be >= before")
	assert.True(t, !df.CreatedAt.After(after), "CreatedAt should be <= after")
}

func TestDateFieldsSaving(t *testing.T) {
	df := &DateFields{}

	before := time.Now().UTC()
	err := df.Saving()
	after := time.Now().UTC()

	require.NoError(t, err)
	assert.False(t, df.UpdatedAt.IsZero(), "UpdatedAt should be set")
	assert.True(t, !df.UpdatedAt.Before(before), "UpdatedAt should be >= before")
	assert.True(t, !df.UpdatedAt.After(after), "UpdatedAt should be <= after")
}

func TestDefaultModelCreatingSetsCreatedAt(t *testing.T) {
	model := &DefaultModel{}

	before := time.Now().UTC()
	err := model.Creating()
	after := time.Now().UTC()

	require.NoError(t, err)
	assert.False(t, model.DateFields.CreatedAt.IsZero())
	assert.True(t, !model.DateFields.CreatedAt.Before(before))
	assert.True(t, !model.DateFields.CreatedAt.After(after))
}

func TestDefaultModelSavingSetsUpdatedAt(t *testing.T) {
	model := &DefaultModel{}

	before := time.Now().UTC()
	err := model.Saving()
	after := time.Now().UTC()

	require.NoError(t, err)
	assert.False(t, model.DateFields.UpdatedAt.IsZero())
	assert.True(t, !model.DateFields.UpdatedAt.Before(before))
	assert.True(t, !model.DateFields.UpdatedAt.After(after))
}

func TestIDFieldGetID(t *testing.T) {
	id := bson.NewObjectID()
	f := &IDField{ID: id}

	got := f.GetID()
	assert.Equal(t, id, got)
}

func TestIDFieldGetIDZeroValue(t *testing.T) {
	f := &IDField{}

	got := f.GetID()
	assert.Equal(t, bson.ObjectID{}, got)
}

func TestIDFieldSetID(t *testing.T) {
	f := &IDField{}
	id := bson.NewObjectID()

	f.SetID(id)
	assert.Equal(t, id, f.ID)
}

func TestIDFieldSetIDPanicsOnWrongType(t *testing.T) {
	f := &IDField{}

	assert.Panics(t, func() {
		f.SetID("not-an-object-id")
	}, "SetID with wrong type should panic")
}

func TestIDFieldSetIDPanicsOnNil(t *testing.T) {
	f := &IDField{}

	assert.Panics(t, func() {
		f.SetID(nil)
	}, "SetID with nil should panic")
}

func TestPrepareIDWithObjectID(t *testing.T) {
	f := &IDField{}
	id := bson.NewObjectID()

	result, err := f.PrepareID(id)

	require.NoError(t, err)
	assert.Equal(t, id, result)
}

func TestPrepareIDWithValidHex(t *testing.T) {
	f := &IDField{}
	hexID := "5df7fb2b1fff9ee374b6bd2a"

	result, err := f.PrepareID(hexID)

	require.NoError(t, err)
	expected, _ := bson.ObjectIDFromHex(hexID)
	assert.Equal(t, expected, result)
}

func TestPrepareIDWithEmptyString(t *testing.T) {
	f := &IDField{}

	_, err := f.PrepareID("")
	assert.Error(t, err, "Empty string should return error")
}

func TestPrepareIDWithShortHex(t *testing.T) {
	f := &IDField{}

	_, err := f.PrepareID("abc")
	assert.Error(t, err, "Short hex string should return error")
}

func TestPrepareIDWithNonStringNonObjectID(t *testing.T) {
	f := &IDField{}

	// Non-string values pass through unchanged
	result, err := f.PrepareID(12345)

	require.NoError(t, err)
	assert.Equal(t, 12345, result)
}

func TestPrepareIDWithNil(t *testing.T) {
	f := &IDField{}

	result, err := f.PrepareID(nil)

	require.NoError(t, err)
	assert.Nil(t, result)
}
