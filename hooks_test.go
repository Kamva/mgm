package mgm_test

import (
	"context"
	"errors"
	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/internal/util"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

func insertPerson(person *Person, ctx context.Context) {
	// Set listeners to mocked hooks:
	person.On("Creating", ctx).Return(nil)
	person.On("Created", ctx).Return(nil)
	person.On("Saving", ctx).Return(nil)
	person.On("Saved", ctx).Return(nil)

	util.PanicErr(mgm.Coll(person).CreateWithCtx(ctx, person))
}

//--------------------------------
// Mock hooks
//--------------------------------

func (d *Person) Creating(ctx context.Context) error {
	args := d.Called(ctx)
	return args.Error(0)
}

func (d *Person) Created(ctx context.Context) error {
	args := d.Called(ctx)
	return args.Error(0)
}

func (d *Person) Updating(ctx context.Context) error {
	args := d.Called(ctx)
	return args.Error(0)
}

func (d *Person) Updated(ctx context.Context, result *mongo.UpdateResult) error {
	args := d.Called(ctx, result.MatchedCount, result.ModifiedCount)
	return args.Error(0)
}

func (d *Person) Saving(ctx context.Context) error {
	args := d.Called(ctx)
	return args.Error(0)
}

func (d *Person) Saved(ctx context.Context) error {
	args := d.Called(ctx)
	return args.Error(0)
}

func (d *Person) Deleting(ctx context.Context) error {
	args := d.Called(ctx)
	return args.Error(0)
}

func (d *Person) Deleted(ctx context.Context, result *mongo.DeleteResult) error {
	args := d.Called(ctx, result.DeletedCount)
	return args.Error(0)
}

func TestReturnErrorInCreatingHook(t *testing.T) {
	setupDefConnection()
	resetCollection()

	creatingErr := errors.New("test error")
	person := NewPerson("Ali", 24)

	// Set listeners to mocked hooks:
	ctx := mgm.Ctx()
	person.On("Creating", ctx).Return(creatingErr)

	err := mgm.Coll(person).CreateWithCtx(ctx, person)

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
	ctx := mgm.Ctx()
	person.On("Creating", ctx).Return(nil)
	person.On("Created", ctx).Return(nil)
	person.On("Saving", ctx).Return(nil)
	person.On("Saved", ctx).Return(nil)

	util.AssertErrIsNil(t, mgm.Coll(person).CreateWithCtx(ctx, person))
	person.AssertExpectations(t)
}

func TestReturnErrorInSavingHook(t *testing.T) {
	setupDefConnection()
	resetCollection()

	savingErr := errors.New("test error")
	person := NewPerson("Ali", 24)

	// Set listeners to mocked hooks:
	ctx := mgm.Ctx()
	person.On("Creating", ctx).Return(nil)
	person.On("Saving", ctx).Return(savingErr)

	err := mgm.Coll(person).CreateWithCtx(ctx, person)

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
	ctx := mgm.Ctx()
	person.On("Creating", ctx).Return(nil)
	person.On("Created", ctx).Return(nil)
	person.On("Saving", ctx).Return(nil)
	person.On("Saved", ctx).Return(nil)

	util.AssertErrIsNil(t, mgm.Coll(person).CreateWithCtx(ctx, person))
	person.AssertExpectations(t)
}
func TestReturnErrorInUpdatingHook(t *testing.T) {
	setupDefConnection()
	resetCollection()
	ctx := mgm.Ctx()

	oldName := "Ali"
	updatingErr := errors.New("test error")
	person := NewPerson(oldName, 24)

	insertPerson(person, ctx)
	person.Name = "Mehran"

	// Set listeners to mocked hooks:
	person.On("Updating", ctx).Return(updatingErr)

	err := mgm.Coll(person).UpdateWithCtx(ctx, person)

	require.Equal(t, updatingErr, err, "Expected returning hook's error")
	person.AssertExpectations(t)

	// Expected do not update this model:
	oldPerson := &Person{}
	util.PanicErr(mgm.Coll(person).FindByID(person.ID, oldPerson))
	require.Equal(t, oldName, oldPerson.Name, "Expected person's name be %s name, but is %s", oldName, person.Name)
}

func TestUpdatingDocHooks(t *testing.T) {
	setupDefConnection()
	resetCollection()
	ctx := mgm.Ctx()

	newName := "Mehran"
	person := NewPerson("Ali", 24)

	insertPerson(person, ctx)

	person.Name = newName

	// Set listeners to mocked hooks:
	person.On("Updating", ctx).Return(nil)
	person.On("Updated", ctx, int64(1), int64(1)).Return(nil)

	err := mgm.Coll(person).UpdateWithCtx(ctx, person)

	util.AssertErrIsNil(t, err)
	person.AssertExpectations(t)

	// Expected do not update this model:
	newPerson := &Person{}
	util.PanicErr(mgm.Coll(person).FindByID(person.ID, newPerson))
	require.Equal(t, newName, newPerson.Name, "Expected person's name be %s , but is %s", newName, person.Name)
}

func TestReturnErrorInDeletingHook(t *testing.T) {
	setupDefConnection()
	resetCollection()
	ctx := mgm.Ctx()

	deletingErr := errors.New("test error")
	person := NewPerson("Ali", 24)

	insertPerson(person, ctx)

	// Set listeners to mocked hooks:
	person.On("Deleting", ctx).Return(deletingErr)

	err := mgm.Coll(person).DeleteWithCtx(ctx, person)

	require.Equal(t, deletingErr, err, "Expected returning hook's error")
	person.AssertExpectations(t)

	// Expected do not delete this model:
	count, _ := mgm.Coll(person).CountDocuments(mgm.Ctx(), bson.M{})
	require.Equal(t, count, int64(1), "Expected having one document,got ", count)
}

func TestDeletingDocHooks(t *testing.T) {
	setupDefConnection()
	resetCollection()
	ctx := mgm.Ctx()

	person := NewPerson("Ali", 24)

	insertPerson(person, ctx)

	// Set listeners to mocked hooks:
	person.On("Deleting", ctx).Return(nil)
	person.On("Deleted", ctx, int64(1)).Return(nil)

	util.AssertErrIsNil(t, mgm.Coll(person).DeleteWithCtx(ctx, person))

	person.AssertExpectations(t)

	// Expected do not delete this model:
	count, _ := mgm.Coll(person).CountDocuments(mgm.Ctx(), bson.M{})
	require.Equal(t, count, int64(0), "Expected having no documents,got ", count)
}
