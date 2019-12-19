package mgm_test

import (
	"github.com/stretchr/testify/require"
	"mgm"
	"mgm/internal"
	"testing"
)

func TestFindDocWithInvalidId(t *testing.T) {
	setupDefConnection()
	resetCollection()
	seed()

	require.NotNil(t, mgm.ModelCollection(&Doc{}).First("invalid id", &Doc{}))
}

func TestCreateDoc(t *testing.T) {
	setupDefConnection()
	resetCollection()

	doc := NewDoc("Ali", 24)

	internal.AssertErrIsNil(t, doc.Collection().Create(doc))

	// Inserted model's id should not be nil:
	require.NotNil(t, doc.Id, "Expected document having id after insertion, got nil")

	// We should have one document in database that is equal to this doc:
	foundDoc := &Doc{}
	internal.AssertErrIsNil(t, doc.Collection().First(doc.Id, foundDoc))

	require.Equal(t, doc.Name, foundDoc.Name, "expected inserted and retrieved docs be equal, got %v and %v", doc.Name, foundDoc.Name)
	require.Equal(t, doc.Age, foundDoc.Age, "expected inserted and retrieved docs be equal, got %v and %v", doc.Age, foundDoc.Age)
}

func TestSaveNewDoc(t *testing.T) {
	setupDefConnection()
	resetCollection()

	doc := NewDoc("Ali", 24)

	internal.AssertErrIsNil(t, doc.Collection().Save(doc))

	// Inserted model's id should not be nil:
	require.NotNil(t, doc.Id, "Expected document having id after save, got nil")

	// We should have one document in database that is equal to this doc:
	foundDoc := &Doc{}
	internal.AssertErrIsNil(t, doc.Collection().First(doc.Id, foundDoc))

	require.Equal(t, doc.Name, foundDoc.Name, "expected inserted and retrieved docs be equal, got %v and %v", doc.Name, foundDoc.Name)
	require.Equal(t, doc.Age, foundDoc.Age, "expected inserted and retrieved docs be equal, got %v and %v", doc.Age, foundDoc.Age)
}

func TestUpdateDoc(t *testing.T) {
	setupDefConnection()
	resetCollection()
	seed()

	found := findDoc(t)

	found.Name = found.Name + "_extra_val"
	found.Age = found.Age + 4

	internal.AssertErrIsNil(t, found.Collection().Update(found))

	// Find that doc again:
	newFound := findDoc(t)

	if found.Id != newFound.Id {
		panic("two fond document dont have same id!")
	}
	require.Equal(t, found.Name, newFound.Name)
	require.Equal(t, found.Age, newFound.Age)
}

func TestSaveExistedDoc(t *testing.T) {
	setupDefConnection()
	resetCollection()
	seed()

	found := findDoc(t)

	found.Name = found.Name + "_extra_val"
	found.Age = found.Age + 4

	internal.AssertErrIsNil(t, found.Collection().Save(found))

	// Find that doc again:
	newFound := findDoc(t)

	if found.Id != newFound.Id {
		panic("two fond document dont have same id!")
	}

	require.Equal(t, found.Name, newFound.Name)
	require.Equal(t, found.Age, newFound.Age)
}

func TestDeleteDoc(t *testing.T) {
	setupDefConnection()
	resetCollection()
	seed()

	found := findDoc(t)

	internal.AssertErrIsNil(t, found.Collection().Delete(found))

	// Find that doc again:
	newFound := findDoc(t)

	require.NotEqual(t, found.Id, newFound.Id)
}
