package mgm

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// testModel is a minimal Model for testing hooks without DB.
type testModel struct {
	IDField    `bson:",inline"`
	DateFields `bson:",inline"`
}

// --- Models with only legacy hooks (no context) ---

type legacyHookModel struct {
	testModel
	creatingCalled bool
	savingCalled   bool
	createdCalled  bool
	savedCalled    bool
	updatingCalled bool
	updatedCalled  bool
	deletingCalled bool
	deletedCalled  bool
}

func (m *legacyHookModel) Creating() error {
	m.creatingCalled = true
	return nil
}

func (m *legacyHookModel) Saving() error {
	m.savingCalled = true
	return nil
}

func (m *legacyHookModel) Created() error {
	m.createdCalled = true
	return nil
}

func (m *legacyHookModel) Saved() error {
	m.savedCalled = true
	return nil
}

func (m *legacyHookModel) Updating() error {
	m.updatingCalled = true
	return nil
}

func (m *legacyHookModel) Updated(result *mongo.UpdateResult) error {
	m.updatedCalled = true
	return nil
}

func (m *legacyHookModel) Deleting() error {
	m.deletingCalled = true
	return nil
}

func (m *legacyHookModel) Deleted(result *mongo.DeleteResult) error {
	m.deletedCalled = true
	return nil
}

// --- Models with context hooks ---

type ctxHookModel struct {
	testModel
	creatingCtx context.Context
	savingCtx   context.Context
	createdCtx  context.Context
	savedCtx    context.Context
	updatingCtx context.Context
	updatedCtx  context.Context
	deletingCtx context.Context
	deletedCtx  context.Context
}

func (m *ctxHookModel) Creating(ctx context.Context) error {
	m.creatingCtx = ctx
	return nil
}

func (m *ctxHookModel) Saving(ctx context.Context) error {
	m.savingCtx = ctx
	return nil
}

func (m *ctxHookModel) Created(ctx context.Context) error {
	m.createdCtx = ctx
	return nil
}

func (m *ctxHookModel) Saved(ctx context.Context) error {
	m.savedCtx = ctx
	return nil
}

func (m *ctxHookModel) Updating(ctx context.Context) error {
	m.updatingCtx = ctx
	return nil
}

func (m *ctxHookModel) Updated(ctx context.Context, result *mongo.UpdateResult) error {
	m.updatedCtx = ctx
	return nil
}

func (m *ctxHookModel) Deleting(ctx context.Context) error {
	m.deletingCtx = ctx
	return nil
}

func (m *ctxHookModel) Deleted(ctx context.Context, result *mongo.DeleteResult) error {
	m.deletedCtx = ctx
	return nil
}

// --- Model with no hooks at all ---

type noHookModel struct {
	testModel
}

// --- Error-returning hook model ---

type errorHookModel struct {
	testModel
	creatingErr error
	savingErr   error
}

func (m *errorHookModel) Creating() error {
	return m.creatingErr
}

func (m *errorHookModel) Saving() error {
	return m.savingErr
}

// ------ Tests ------

func TestCallToBeforeCreateHooks_LegacyFallback(t *testing.T) {
	model := &legacyHookModel{}
	model.ID = bson.NewObjectID()

	err := callToBeforeCreateHooks(context.Background(), model)

	require.NoError(t, err)
	assert.True(t, model.creatingCalled, "Legacy Creating() should be called")
	assert.True(t, model.savingCalled, "Legacy Saving() should be called")
}

func TestCallToBeforeCreateHooks_CtxPreferred(t *testing.T) {
	model := &ctxHookModel{}
	model.ID = bson.NewObjectID()
	ctx := context.WithValue(context.Background(), "key", "value")

	err := callToBeforeCreateHooks(ctx, model)

	require.NoError(t, err)
	assert.Equal(t, ctx, model.creatingCtx, "Context should be passed to Creating(ctx)")
	assert.Equal(t, ctx, model.savingCtx, "Context should be passed to Saving(ctx)")
}

func TestCallToBeforeCreateHooks_NoHooks(t *testing.T) {
	model := &noHookModel{}
	model.ID = bson.NewObjectID()

	err := callToBeforeCreateHooks(context.Background(), model)
	require.NoError(t, err)
}

func TestCallToAfterCreateHooks_LegacyFallback(t *testing.T) {
	model := &legacyHookModel{}
	model.ID = bson.NewObjectID()

	err := callToAfterCreateHooks(context.Background(), model)

	require.NoError(t, err)
	assert.True(t, model.createdCalled, "Legacy Created() should be called")
	assert.True(t, model.savedCalled, "Legacy Saved() should be called")
}

func TestCallToAfterCreateHooks_CtxPreferred(t *testing.T) {
	model := &ctxHookModel{}
	model.ID = bson.NewObjectID()
	ctx := context.WithValue(context.Background(), "key", "value")

	err := callToAfterCreateHooks(ctx, model)

	require.NoError(t, err)
	assert.Equal(t, ctx, model.createdCtx)
	assert.Equal(t, ctx, model.savedCtx)
}

func TestCallToBeforeUpdateHooks_LegacyFallback(t *testing.T) {
	model := &legacyHookModel{}
	model.ID = bson.NewObjectID()

	err := callToBeforeUpdateHooks(context.Background(), model)

	require.NoError(t, err)
	assert.True(t, model.updatingCalled)
	assert.True(t, model.savingCalled)
}

func TestCallToBeforeUpdateHooks_CtxPreferred(t *testing.T) {
	model := &ctxHookModel{}
	model.ID = bson.NewObjectID()
	ctx := context.WithValue(context.Background(), "key", "value")

	err := callToBeforeUpdateHooks(ctx, model)

	require.NoError(t, err)
	assert.Equal(t, ctx, model.updatingCtx)
	assert.Equal(t, ctx, model.savingCtx)
}

func TestCallToAfterUpdateHooks_LegacyFallback(t *testing.T) {
	model := &legacyHookModel{}
	model.ID = bson.NewObjectID()
	result := &mongo.UpdateResult{MatchedCount: 1, ModifiedCount: 1}

	err := callToAfterUpdateHooks(context.Background(), result, model)

	require.NoError(t, err)
	assert.True(t, model.updatedCalled)
	assert.True(t, model.savedCalled)
}

func TestCallToAfterUpdateHooks_CtxPreferred(t *testing.T) {
	model := &ctxHookModel{}
	model.ID = bson.NewObjectID()
	ctx := context.WithValue(context.Background(), "key", "value")
	result := &mongo.UpdateResult{MatchedCount: 1, ModifiedCount: 1}

	err := callToAfterUpdateHooks(ctx, result, model)

	require.NoError(t, err)
	assert.Equal(t, ctx, model.updatedCtx)
	assert.Equal(t, ctx, model.savedCtx)
}

func TestCallToBeforeDeleteHooks_LegacyFallback(t *testing.T) {
	model := &legacyHookModel{}
	model.ID = bson.NewObjectID()

	err := callToBeforeDeleteHooks(context.Background(), model)

	require.NoError(t, err)
	assert.True(t, model.deletingCalled)
}

func TestCallToBeforeDeleteHooks_CtxPreferred(t *testing.T) {
	model := &ctxHookModel{}
	model.ID = bson.NewObjectID()
	ctx := context.WithValue(context.Background(), "key", "value")

	err := callToBeforeDeleteHooks(ctx, model)

	require.NoError(t, err)
	assert.Equal(t, ctx, model.deletingCtx)
}

func TestCallToAfterDeleteHooks_LegacyFallback(t *testing.T) {
	model := &legacyHookModel{}
	model.ID = bson.NewObjectID()
	result := &mongo.DeleteResult{DeletedCount: 1}

	err := callToAfterDeleteHooks(context.Background(), result, model)

	require.NoError(t, err)
	assert.True(t, model.deletedCalled)
}

func TestCallToAfterDeleteHooks_CtxPreferred(t *testing.T) {
	model := &ctxHookModel{}
	model.ID = bson.NewObjectID()
	ctx := context.WithValue(context.Background(), "key", "value")
	result := &mongo.DeleteResult{DeletedCount: 1}

	err := callToAfterDeleteHooks(ctx, result, model)

	require.NoError(t, err)
	assert.Equal(t, ctx, model.deletedCtx)
}

func TestCallToBeforeCreateHooks_CreatingErrorShortCircuits(t *testing.T) {
	expectedErr := errors.New("creating error")
	model := &errorHookModel{creatingErr: expectedErr}
	model.ID = bson.NewObjectID()

	err := callToBeforeCreateHooks(context.Background(), model)

	require.Equal(t, expectedErr, err, "Creating error should propagate")
}

func TestCallToBeforeCreateHooks_SavingErrorPropagates(t *testing.T) {
	expectedErr := errors.New("saving error")
	model := &errorHookModel{savingErr: expectedErr}
	model.ID = bson.NewObjectID()

	err := callToBeforeCreateHooks(context.Background(), model)

	require.Equal(t, expectedErr, err, "Saving error should propagate")
}

func TestCallToBeforeCreateHooks_CreatingErrorSkipsSaving(t *testing.T) {
	// When Creating returns error, Saving should not be called.
	// We can verify this indirectly: if Creating errors, we get that error back,
	// not the saving error.
	creatingErr := errors.New("creating error")
	savingErr := errors.New("saving error")
	model := &errorHookModel{creatingErr: creatingErr, savingErr: savingErr}
	model.ID = bson.NewObjectID()

	err := callToBeforeCreateHooks(context.Background(), model)

	require.Equal(t, creatingErr, err, "Should get creating error, not saving error")
}
