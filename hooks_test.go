package mgm_test

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"mgm"
	"mgm/internal"
	"testing"
)

type Person struct {
	mock.Mock        `bson:"-"`
	mgm.DefaultModel `bson:",inline"`

	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

func (d *Person) CollectionName() string {
	return "persons"
}

func NewPerson(name string, age int) *Person {
	return &Person{Name: name, Age: age}
}

func insertPerson(person *Person) {
	// Set listeners to mocked hooks:
	person.On("Creating").Return(nil)
	person.On("Created").Return(nil)
	person.On("Saving").Return(nil)
	person.On("Saved").Return(nil)

	internal.PanicErr(mgm.Coll(person).Save(person))
}

//--------------------------------
// Mock hooks
//--------------------------------

func (d *Person) Creating() error {
	args := d.Called()
	return args.Error(0)
}

func (d *Person) Created() error {
	args := d.Called()
	return args.Error(0)
}

func (d *Person) Updating() error {
	args := d.Called()
	return args.Error(0)
}

func (d *Person) Updated(result *mongo.UpdateResult) error {
	args := d.Called(result.MatchedCount, result.ModifiedCount)
	return args.Error(0)
}

func (d *Person) Saving() error {
	args := d.Called()
	return args.Error(0)
}

func (d *Person) Saved() error {
	args := d.Called()
	return args.Error(0)
}

func (d *Person) Deleting() error {
	args := d.Called()
	return args.Error(0)
}

func (d *Person) Deleted(result *mongo.DeleteResult) error {
	args := d.Called(result.DeletedCount)
	return args.Error(0)
}

func TestReturnErrorInCreatingHook(t *testing.T) {
	setupDefConnection()
	resetCollection()

	creatingErr := errors.New("test error")
	person := NewPerson("Ali", 24)

	// Set listeners to mocked hooks:
	person.On("Creating").Return(creatingErr)

	err := mgm.Coll(person).Create(person)

	require.Equal(t, creatingErr, err, "Expected returning hook's error")
	person.AssertExpectations(t)

	// Expected do not inserting this model:
	count, _ := mgm.Coll(person).CountDocuments(mgm.Ctx(), bson.M{})
	require.Equal(t, count, int64(0))
}

func TestCreatingDocHooks(t *testing.T) {
	setupDefConnection()
	resetCollection()

	person := NewPerson("Ali", 24)

	// Set listeners to mocked hooks:
	person.On("Creating").Return(nil)
	person.On("Created").Return(nil)
	person.On("Saving").Return(nil)
	person.On("Saved").Return(nil)

	internal.AssertErrIsNil(t, mgm.Coll(person).Create(person))
	person.AssertExpectations(t)
}

func TestReturnErrorInSavingHook(t *testing.T) {
	setupDefConnection()
	resetCollection()

	savingErr := errors.New("test error")
	person := NewPerson("Ali", 24)

	// Set listeners to mocked hooks:
	person.On("Creating").Return(nil)
	person.On("Saving").Return(savingErr)

	err := mgm.Coll(person).Save(person)

	require.Equal(t, savingErr, err, "Expected returning hook's error")
	person.AssertExpectations(t)

	// Expected do not inserting this model:
	count, _ := mgm.Coll(person).CountDocuments(mgm.Ctx(), bson.M{})
	require.Equal(t, count, int64(0))
}

func TestSavingDocHooks(t *testing.T) {
	setupDefConnection()
	resetCollection()

	person := NewPerson("Ali", 24)

	// Set listeners to mocked hooks:
	person.On("Creating").Return(nil)
	person.On("Created").Return(nil)
	person.On("Saving").Return(nil)
	person.On("Saved").Return(nil)

	internal.AssertErrIsNil(t, mgm.Coll(person).Create(person))
	person.AssertExpectations(t)
}
func TestReturnErrorInUpdatingHook(t *testing.T) {
	setupDefConnection()
	resetCollection()
	oldName := "Ali"
	updatingErr := errors.New("test error")
	person := NewPerson(oldName, 24)

	insertPerson(person)
	person.Name = "Mehran"

	// Set listeners to mocked hooks:
	person.On("Updating").Return(updatingErr)

	err := mgm.Coll(person).Update(person)

	require.Equal(t, updatingErr, err, "Expected returning hook's error")
	person.AssertExpectations(t)

	// Expected do not update this model:
	oldPerson := &Person{}
	internal.PanicErr(mgm.Coll(person).FindById(person.Id, oldPerson))
	require.Equal(t, oldName, oldPerson.Name, "Expected person's name be %s name, but is %s", oldName, person.Name)
}

func TestUpdatingDocHooks(t *testing.T) {
	setupDefConnection()
	resetCollection()
	newName := "Mehran"
	person := NewPerson("Ali", 24)

	insertPerson(person)

	person.Name = newName

	// Set listeners to mocked hooks:
	person.On("Updating").Return(nil)
	person.On("Updated", int64(1), int64(1)).Return(nil)

	err := mgm.Coll(person).Update(person)

	internal.AssertErrIsNil(t, err)
	person.AssertExpectations(t)

	// Expected do not update this model:
	newPerson := &Person{}
	internal.PanicErr(mgm.Coll(person).FindById(person.Id, newPerson))
	require.Equal(t, newName, newPerson.Name, "Expected person's name be %s , but is %s", newName, person.Name)
}

func TestReturnErrorInDeletingHook(t *testing.T) {
	setupDefConnection()
	resetCollection()
	deletingErr := errors.New("test error")
	person := NewPerson("Ali", 24)

	insertPerson(person)

	// Set listeners to mocked hooks:
	person.On("Deleting").Return(deletingErr)

	err := mgm.Coll(person).Delete(person)

	require.Equal(t, deletingErr, err, "Expected returning hook's error")
	person.AssertExpectations(t)

	// Expected do not delete this model:
	count, _ := mgm.Coll(person).CountDocuments(mgm.Ctx(), bson.M{})
	require.Equal(t, count, int64(1), "Expected having one document,got ", count)
}

func TestDeletingDocHooks(t *testing.T) {
	setupDefConnection()
	resetCollection()
	person := NewPerson("Ali", 24)

	insertPerson(person)

	// Set listeners to mocked hooks:
	person.On("Deleting").Return(nil)
	person.On("Deleted", int64(1)).Return(nil)

	internal.AssertErrIsNil(t, mgm.Coll(person).Delete(person))

	person.AssertExpectations(t)

	// Expected do not delete this model:
	count, _ := mgm.Coll(person).CountDocuments(mgm.Ctx(), bson.M{})
	require.Equal(t, count, int64(0), "Expected having no documents,got ", count)
}
