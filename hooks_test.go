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

func insertPerson(person *Person) {
	// Set listeners to mocked hooks:
	person.On("Creating").Return(nil)
	person.On("Created").Return(nil)
	person.On("Saving").Return(nil)
	person.On("Saved").Return(nil)

	util.PanicErr(mgm.Coll(person).Create(person))
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

	util.AssertErrIsNil(t, mgm.Coll(person).Create(person))
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

	err := mgm.Coll(person).Create(person)

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

	util.AssertErrIsNil(t, mgm.Coll(person).Create(person))
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
	util.PanicErr(mgm.Coll(person).FindByID(person.ID, oldPerson))
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

	util.AssertErrIsNil(t, mgm.Coll(person).Delete(person))

	person.AssertExpectations(t)

	// Expected do not delete this model:
	count, _ := mgm.Coll(person).CountDocuments(mgm.Ctx(), bson.M{})
	require.Equal(t, count, int64(0), "Expected having no documents,got ", count)
}

type Celebrity struct {
	mock.Mock        `bson:"-"`
	mgm.DefaultModel `bson:",inline"`

	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

func (d *Celebrity) CollectionName() string {
	return "persons"
}

func NewCelebrity(name string, age int) *Celebrity {
	return &Celebrity{Name: name, Age: age}
}

func insertCelebrity(person *Celebrity, ctx context.Context) {
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

func (d *Celebrity) Creating(ctx context.Context) error {
	args := d.Called(ctx)
	return args.Error(0)
}

func (d *Celebrity) Created(ctx context.Context) error {
	args := d.Called(ctx)
	return args.Error(0)
}

func (d *Celebrity) Updating(ctx context.Context) error {
	args := d.Called(ctx)
	return args.Error(0)
}

func (d *Celebrity) Updated(ctx context.Context, result *mongo.UpdateResult) error {
	args := d.Called(ctx, result.MatchedCount, result.ModifiedCount)
	return args.Error(0)
}

func (d *Celebrity) Saving(ctx context.Context) error {
	args := d.Called(ctx)
	return args.Error(0)
}

func (d *Celebrity) Saved(ctx context.Context) error {
	args := d.Called(ctx)
	return args.Error(0)
}

func (d *Celebrity) Deleting(ctx context.Context) error {
	args := d.Called(ctx)
	return args.Error(0)
}

func (d *Celebrity) Deleted(ctx context.Context, result *mongo.DeleteResult) error {
	args := d.Called(ctx, result.DeletedCount)
	return args.Error(0)
}

func TestReturnErrorInCreatingHook_Celebrity(t *testing.T) {
	setupDefConnection()
	resetCollection()

	creatingErr := errors.New("test error")
	celebrity := NewCelebrity("Ali", 24)

	// Set listeners to mocked hooks:
	ctx := mgm.Ctx()
	celebrity.On("Creating", ctx).Return(creatingErr)

	err := mgm.Coll(celebrity).CreateWithCtx(ctx, celebrity)

	require.Equal(t, creatingErr, err, "Expected returning hook's error")
	celebrity.AssertExpectations(t)

	// Expected do not inserting this model:
	count, _ := mgm.Coll(celebrity).CountDocuments(mgm.Ctx(), bson.M{})
	require.Equal(t, count, int64(0))
}

func TestCreatingDocHooks_Celebrity(t *testing.T) {
	setupDefConnection()
	resetCollection()

	celebrity := NewCelebrity("Ali", 24)

	// Set listeners to mocked hooks:
	ctx := mgm.Ctx()
	celebrity.On("Creating", ctx).Return(nil)
	celebrity.On("Created", ctx).Return(nil)
	celebrity.On("Saving", ctx).Return(nil)
	celebrity.On("Saved", ctx).Return(nil)

	util.AssertErrIsNil(t, mgm.Coll(celebrity).CreateWithCtx(ctx, celebrity))
	celebrity.AssertExpectations(t)
}

func TestReturnErrorInSavingHook_Celebrity(t *testing.T) {
	setupDefConnection()
	resetCollection()

	savingErr := errors.New("test error")
	celebrity := NewCelebrity("Ali", 24)

	// Set listeners to mocked hooks:
	ctx := mgm.Ctx()
	celebrity.On("Creating", ctx).Return(nil)
	celebrity.On("Saving", ctx).Return(savingErr)

	err := mgm.Coll(celebrity).CreateWithCtx(ctx, celebrity)

	require.Equal(t, savingErr, err, "Expected returning hook's error")
	celebrity.AssertExpectations(t)

	// Expected do not inserting this model:
	count, _ := mgm.Coll(celebrity).CountDocuments(mgm.Ctx(), bson.M{})
	require.Equal(t, count, int64(0))
}

func TestSavingDocHooks_Celebrity(t *testing.T) {
	setupDefConnection()
	resetCollection()

	celebrity := NewCelebrity("Ali", 24)

	// Set listeners to mocked hooks:
	ctx := mgm.Ctx()
	celebrity.On("Creating", ctx).Return(nil)
	celebrity.On("Created", ctx).Return(nil)
	celebrity.On("Saving", ctx).Return(nil)
	celebrity.On("Saved", ctx).Return(nil)

	util.AssertErrIsNil(t, mgm.Coll(celebrity).CreateWithCtx(ctx, celebrity))
	celebrity.AssertExpectations(t)
}
func TestReturnErrorInUpdatingHook_Celebrity(t *testing.T) {
	setupDefConnection()
	resetCollection()
	ctx := mgm.Ctx()

	oldName := "Ali"
	updatingErr := errors.New("test error")
	celebrity := NewCelebrity(oldName, 24)

	insertCelebrity(celebrity, ctx)
	celebrity.Name = "Mehran"

	// Set listeners to mocked hooks:
	celebrity.On("Updating", ctx).Return(updatingErr)

	err := mgm.Coll(celebrity).UpdateWithCtx(ctx, celebrity)

	require.Equal(t, updatingErr, err, "Expected returning hook's error")
	celebrity.AssertExpectations(t)

	// Expected do not update this model:
	oldPerson := &Person{}
	util.PanicErr(mgm.Coll(celebrity).FindByID(celebrity.ID, oldPerson))
	require.Equal(t, oldName, oldPerson.Name, "Expected celebrity's name be %s name, but is %s", oldName, celebrity.Name)
}

func TestUpdatingDocHooks_Celebrity(t *testing.T) {
	setupDefConnection()
	resetCollection()
	ctx := mgm.Ctx()

	newName := "Mehran"
	celebrity := NewCelebrity("Ali", 24)

	insertCelebrity(celebrity, ctx)

	celebrity.Name = newName

	// Set listeners to mocked hooks:
	celebrity.On("Updating", ctx).Return(nil)
	celebrity.On("Updated", ctx, int64(1), int64(1)).Return(nil)

	err := mgm.Coll(celebrity).UpdateWithCtx(ctx, celebrity)

	util.AssertErrIsNil(t, err)
	celebrity.AssertExpectations(t)

	// Expected do not update this model:
	newCelebrity := &Celebrity{}
	util.PanicErr(mgm.Coll(celebrity).FindByID(celebrity.ID, newCelebrity))
	require.Equal(t, newName, newCelebrity.Name, "Expected celebrity's name be %s , but is %s", newName, celebrity.Name)
}

func TestReturnErrorInDeletingHook_Celebrity(t *testing.T) {
	setupDefConnection()
	resetCollection()
	ctx := mgm.Ctx()

	deletingErr := errors.New("test error")
	celebrity := NewCelebrity("Ali", 24)

	insertCelebrity(celebrity, ctx)

	// Set listeners to mocked hooks:
	celebrity.On("Deleting", ctx).Return(deletingErr)

	err := mgm.Coll(celebrity).DeleteWithCtx(ctx, celebrity)

	require.Equal(t, deletingErr, err, "Expected returning hook's error")
	celebrity.AssertExpectations(t)

	// Expected do not delete this model:
	count, _ := mgm.Coll(celebrity).CountDocuments(mgm.Ctx(), bson.M{})
	require.Equal(t, count, int64(1), "Expected having one document,got ", count)
}

func TestDeletingDocHooks_Celebrity(t *testing.T) {
	setupDefConnection()
	resetCollection()
	ctx := mgm.Ctx()

	celebrity := NewCelebrity("Ali", 24)

	insertCelebrity(celebrity, ctx)

	// Set listeners to mocked hooks:
	celebrity.On("Deleting", ctx).Return(nil)
	celebrity.On("Deleted", ctx, int64(1)).Return(nil)

	util.AssertErrIsNil(t, mgm.Coll(celebrity).DeleteWithCtx(ctx, celebrity))

	celebrity.AssertExpectations(t)

	// Expected do not delete this model:
	count, _ := mgm.Coll(celebrity).CountDocuments(mgm.Ctx(), bson.M{})
	require.Equal(t, count, int64(0), "Expected having no documents,got ", count)
}
