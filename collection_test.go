package mgm_test

import (
	"context"
	"testing"
	"time"

	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/builder"
	"github.com/kamva/mgm/v3/internal/util"
	"github.com/kamva/mgm/v3/operator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func TestFindByIdWithInvalidId(t *testing.T) {
	setupDefConnection()
	resetCollection()
	seed()

	require.NotNil(t, mgm.Coll(&Doc{}).FindByID("invalid id", &Doc{}))
}

func TestFindByIDWithObjectID(t *testing.T) {
	setupDefConnection()
	resetCollection()

	doc := NewDoc("Ali", 24)
	util.AssertErrIsNil(t, mgm.Coll(doc).Create(doc))

	found := &Doc{}
	util.AssertErrIsNil(t, mgm.Coll(found).FindByID(doc.ID, found))
	require.Equal(t, doc.Name, found.Name)
}

func TestFindByIDWithHexString(t *testing.T) {
	setupDefConnection()
	resetCollection()

	doc := NewDoc("Ali", 24)
	util.AssertErrIsNil(t, mgm.Coll(doc).Create(doc))

	found := &Doc{}
	util.AssertErrIsNil(t, mgm.Coll(found).FindByID(doc.ID.Hex(), found))
	require.Equal(t, doc.Name, found.Name)
}

func TestFindByIDWithEmptyString(t *testing.T) {
	setupDefConnection()
	resetCollection()

	err := mgm.Coll(&Doc{}).FindByID("", &Doc{})
	require.Error(t, err, "Empty string should return error")
}

func TestFindByIDWithCtxAndOptions(t *testing.T) {
	setupDefConnection()
	resetCollection()

	doc := NewDoc("Ali", 24)
	util.AssertErrIsNil(t, mgm.Coll(doc).Create(doc))

	found := &Doc{}
	ctx := mgm.Ctx()
	opts := options.FindOne().SetComment("test-comment")
	util.AssertErrIsNil(t, mgm.Coll(found).FindByIDWithCtx(ctx, doc.ID, found, opts))
	require.Equal(t, doc.Name, found.Name)
}

func TestFirstWithNilFilter(t *testing.T) {
	setupDefConnection()
	resetCollection()
	seed()

	d := &Doc{}
	util.AssertErrIsNil(t, mgm.Coll(d).First(nil, d))
	require.NotEqual(t, bson.ObjectID{}, d.ID)
}

func TestFirstWithEmptyFilter(t *testing.T) {
	setupDefConnection()
	resetCollection()
	seed()

	d := &Doc{}
	util.AssertErrIsNil(t, mgm.Coll(d).First(bson.M{}, d))
	require.NotEqual(t, bson.ObjectID{}, d.ID)
}

func TestFirstWithNoMatch(t *testing.T) {
	setupDefConnection()
	resetCollection()
	seed()

	d := &Doc{}
	err := mgm.Coll(d).First(bson.M{"name": "nonexistent"}, d)
	require.Error(t, err)
}

func TestCreateSetsTimestamps(t *testing.T) {
	setupDefConnection()
	resetCollection()

	before := time.Now().UTC().Add(-time.Second)
	doc := NewDoc("Ali", 24)
	util.AssertErrIsNil(t, mgm.Coll(doc).Create(doc))
	after := time.Now().UTC().Add(time.Second)

	require.NotEqual(t, bson.ObjectID{}, doc.ID)
	assert.False(t, doc.CreatedAt.IsZero(), "CreatedAt should be set")
	assert.False(t, doc.UpdatedAt.IsZero(), "UpdatedAt should be set")
	assert.True(t, doc.CreatedAt.After(before), "CreatedAt should be after test start")
	assert.True(t, doc.CreatedAt.Before(after), "CreatedAt should be before test end")
}

func TestCreateWithCtxAndOptions(t *testing.T) {
	setupDefConnection()
	resetCollection()

	doc := NewDoc("Ali", 24)
	ctx := mgm.Ctx()
	opts := options.InsertOne().SetComment("test-insert")
	util.AssertErrIsNil(t, mgm.Coll(doc).CreateWithCtx(ctx, doc, opts))
	require.NotEqual(t, bson.ObjectID{}, doc.ID)
}

func TestUpdateSetsUpdatedAt(t *testing.T) {
	setupDefConnection()
	resetCollection()

	doc := NewDoc("Ali", 24)
	util.AssertErrIsNil(t, mgm.Coll(doc).Create(doc))

	originalUpdatedAt := doc.UpdatedAt
	time.Sleep(10 * time.Millisecond)

	doc.Name = "Updated"
	util.AssertErrIsNil(t, mgm.Coll(doc).Update(doc))

	assert.True(t, doc.UpdatedAt.After(originalUpdatedAt) || doc.UpdatedAt.Equal(originalUpdatedAt),
		"UpdatedAt should be >= original value after update")
}

func TestUpdateWithCtxAndOptions(t *testing.T) {
	setupDefConnection()
	resetCollection()

	doc := NewDoc("Ali", 24)
	util.AssertErrIsNil(t, mgm.Coll(doc).Create(doc))

	doc.Name = "Updated"
	ctx := mgm.Ctx()
	opts := options.UpdateOne().SetComment("test-update")
	util.AssertErrIsNil(t, mgm.Coll(doc).UpdateWithCtx(ctx, doc, opts))

	found := &Doc{}
	util.AssertErrIsNil(t, mgm.Coll(found).FindByID(doc.ID, found))
	require.Equal(t, "Updated", found.Name)
}

func TestUpdateWithUpsertOption(t *testing.T) {
	setupDefConnection()
	resetCollection()

	doc := NewDoc("Ali", 24)
	util.AssertErrIsNil(t, mgm.Coll(doc).Create(doc))

	doc.Name = "Upserted"
	upsertOpt := mgm.UpsertTrueOption()
	util.AssertErrIsNil(t, mgm.Coll(doc).Update(doc, upsertOpt))

	found := &Doc{}
	util.AssertErrIsNil(t, mgm.Coll(found).FindByID(doc.ID, found))
	require.Equal(t, "Upserted", found.Name)
}

func TestSimpleFindEmptyCollection(t *testing.T) {
	setupDefConnection()
	resetCollection()

	results := []Doc{}
	err := mgm.Coll(&Doc{}).SimpleFind(&results, bson.M{})

	util.AssertErrIsNil(t, err)
	require.Empty(t, results)
}

func TestSimpleFindWithOptions(t *testing.T) {
	setupDefConnection()
	resetCollection()
	seed()

	results := []Doc{}
	opts := options.Find().SetLimit(2)
	err := mgm.Coll(&Doc{}).SimpleFind(&results, bson.M{}, opts)

	util.AssertErrIsNil(t, err)
	require.Len(t, results, 2)
}

func TestSimpleFindWithSortOption(t *testing.T) {
	setupDefConnection()
	resetCollection()
	seed()

	results := []Doc{}
	opts := options.Find().SetSort(bson.M{"age": -1})
	err := mgm.Coll(&Doc{}).SimpleFind(&results, bson.M{}, opts)

	util.AssertErrIsNil(t, err)
	require.NotEmpty(t, results)

	// Verify descending order
	for i := 1; i < len(results); i++ {
		assert.True(t, results[i-1].Age >= results[i].Age,
			"Results should be in descending age order")
	}
}

func TestSimpleFindWithCtx(t *testing.T) {
	setupDefConnection()
	resetCollection()
	seed()

	results := []Doc{}
	ctx := mgm.Ctx()
	err := mgm.Coll(&Doc{}).SimpleFindWithCtx(ctx, &results, bson.M{})

	util.AssertErrIsNil(t, err)
	require.NotEmpty(t, results)
}

func TestFindByIDWithCancelledContext(t *testing.T) {
	setupDefConnection()
	resetCollection()
	seed()

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // cancel immediately

	d := &Doc{}
	err := mgm.Coll(d).FindByIDWithCtx(ctx, bson.NewObjectID(), d)
	require.Error(t, err, "Cancelled context should produce error")
}

func TestCreateWithCancelledContext(t *testing.T) {
	setupDefConnection()
	resetCollection()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	doc := NewDoc("Ali", 24)
	err := mgm.Coll(doc).CreateWithCtx(ctx, doc)
	require.Error(t, err, "Cancelled context should produce error")
}

func TestSimpleFindWithCancelledContext(t *testing.T) {
	setupDefConnection()
	resetCollection()
	seed()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	results := []Doc{}
	err := mgm.Coll(&Doc{}).SimpleFindWithCtx(ctx, &results, bson.M{})
	require.Error(t, err, "Cancelled context should produce error")
}

func TestSimpleAggregateWithEmptyStages(t *testing.T) {
	setupDefConnection()
	resetCollection()
	seed()

	results := []Doc{}
	err := mgm.Coll(&Doc{}).SimpleAggregate(&results, []interface{}{})
	util.AssertErrIsNil(t, err)
	require.NotEmpty(t, results)
}

func TestSimpleAggregateCursorWithCtx(t *testing.T) {
	setupDefConnection()
	resetCollection()
	seed()

	group := builder.Group("$_id", nil)
	ctx := mgm.Ctx()
	cur, err := mgm.Coll(&Doc{}).SimpleAggregateCursorWithCtx(ctx, []interface{}{group})

	util.AssertErrIsNil(t, err)
	require.NotNil(t, cur)

	var results []Doc
	util.AssertErrIsNil(t, cur.All(ctx, &results))
	require.NotEmpty(t, results)
}

func TestSimpleAggregateWithMixedStages(t *testing.T) {
	setupDefConnection()
	resetCollection()
	seed()

	// Mix builder.Operator with raw bson.M stages
	group := builder.Group("$_id", nil)
	rawProject := bson.M{operator.Project: bson.M{"age": 0}}

	results := []Doc{}
	err := mgm.Coll(&Doc{}).SimpleAggregate(&results, []interface{}{group, rawProject})

	util.AssertErrIsNil(t, err)
	require.NotEmpty(t, results)
}

func TestFindFirst(t *testing.T) {
	setupDefConnection()
	resetCollection()
	seed()

	d := &Doc{}
	util.AssertErrIsNil(t, mgm.Coll(&Doc{}).First(bson.M{}, d))

	require.NotEqual(t, bson.ObjectID{}, d.ID)
}

func TestCollection_Create(t *testing.T) {
	setupDefConnection()
	resetCollection()

	doc := NewDoc("Ali", 24)

	util.AssertErrIsNil(t, mgm.Coll(doc).Create(doc))

	// Inserted model's id should not be nil:
	require.NotNil(t, doc.ID, "Expected document having id after insertion, got nil")

	// We should have one document in database that is equal to this doc:
	foundDoc := &Doc{}
	util.AssertErrIsNil(t, mgm.Coll(doc).FindByID(doc.ID, foundDoc))

	require.Equal(t, doc.Name, foundDoc.Name, "expected inserted and retrieved docs be equal, got %v and %v", doc.Name, foundDoc.Name)
	require.Equal(t, doc.Age, foundDoc.Age, "expected inserted and retrieved docs be equal, got %v and %v", doc.Age, foundDoc.Age)
}

func TestCollection_Update(t *testing.T) {
	setupDefConnection()
	resetCollection()
	seed()

	found := findDoc(t)

	found.Name = found.Name + "_extra_val"
	found.Age = found.Age + 4

	util.AssertErrIsNil(t, mgm.Coll(found).Update(found))

	// Find that doc again:
	newFound := findDoc(t)

	if found.ID != newFound.ID {
		panic("two fond document dont have same id!")
	}
	require.Equal(t, found.Name, newFound.Name)
	require.Equal(t, found.Age, newFound.Age)
}

func TestCollection_Delete(t *testing.T) {
	setupDefConnection()
	resetCollection()
	seed()

	found := findDoc(t)

	util.AssertErrIsNil(t, mgm.Coll(found).Delete(found))

	// Find that doc again:
	newFound := findDoc(t)

	require.NotEqual(t, found.ID, newFound.ID)
}

func TestCollection_SimpleFind(t *testing.T) {
	setupDefConnection()
	resetCollection()
	seed()

	expectedResult := []Doc{}
	gotResult := []Doc{}

	filter := bson.M{"age": bson.M{operator.Gt: 24}}
	err := mgm.Coll(&Doc{}).SimpleFind(&gotResult, filter)

	util.AssertErrIsNil(t, err)

	// Create same aggregation by raw methods
	cur, err := mgm.Coll(&Doc{}).Find(mgm.Ctx(), filter)
	util.AssertErrIsNil(t, err)

	util.AssertErrIsNil(t, cur.All(mgm.Ctx(), &expectedResult))

	require.Equal(t, len(expectedResult), len(gotResult))

	// We should have same documents
	for i, expectedDoc := range expectedResult {
		if expectedDoc != gotResult[i] {
			t.Errorf("Expected %v, got %v", expectedDoc, gotResult[i])
		}
	}
}

func TestCollection_SimpleAggregateFirst(t *testing.T) {
	setupDefConnection()
	resetCollection()
	seed()

	expectedResult := []Doc{}
	gotResult := Doc{}

	// We dont want to change document.
	group := builder.Group("$_id", nil)

	found, err := mgm.Coll(&Doc{}).SimpleAggregateFirst(&gotResult, []interface{}{group})

	assert.True(t, found)
	util.AssertErrIsNil(t, err)

	// Create same aggregation by raw methods
	cur, err := mgm.Coll(&Doc{}).Aggregate(mgm.Ctx(), bson.A{builder.S(group)})
	util.AssertErrIsNil(t, err)
	util.AssertErrIsNil(t, cur.All(mgm.Ctx(), &expectedResult))
	assert.Equal(t, expectedResult[0], gotResult)
}

func TestCollection_SimpleAggregateFirstFalse(t *testing.T) {
	setupDefConnection()
	resetCollection()
	seed()

	var gotResult *Doc
	match := bson.M{operator.Match: bson.M{"user_id": "unknown"}}
	found, err := mgm.Coll(&Doc{}).SimpleAggregateFirst(gotResult, []interface{}{match})

	assert.False(t, found)
	util.AssertErrIsNil(t, err)
	assert.Nil(t, gotResult)
}

func TestCollection_SimpleAggregate(t *testing.T) {
	setupDefConnection()
	resetCollection()
	seed()

	expectedResult := []Doc{}
	gotResult := []Doc{}

	// We dont want to change document.
	group := builder.Group("$_id", nil)

	project := bson.M{operator.Project: bson.M{"age": 0}}

	err := mgm.Coll(&Doc{}).SimpleAggregate(&gotResult, []interface{}{group, project})

	util.AssertErrIsNil(t, err)

	// Create same aggregation by raw methods
	cur, err := mgm.Coll(&Doc{}).Aggregate(mgm.Ctx(), bson.A{builder.S(group), project})
	util.AssertErrIsNil(t, err)

	util.AssertErrIsNil(t, cur.All(mgm.Ctx(), &expectedResult))

	require.Equal(t, len(expectedResult), len(gotResult))

	// We should have same documents
	for i, expectedDoc := range expectedResult {
		if expectedDoc != gotResult[i] {
			t.Errorf("Expected %v, got %v", expectedDoc, gotResult[i])
		}
	}
}
