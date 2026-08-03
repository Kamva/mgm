package builder_test

import (
	"testing"

	"github.com/kamva/mgm/v3/builder"
	"github.com/kamva/mgm/v3/operator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestBucketAllNilParams(t *testing.T) {
	res := builder.Bucket(nil, nil, nil, nil)
	require.Equal(t, operator.Bucket, res.GetKey())
	assert.Equal(t, bson.M{}, res.GetVal())
}

func TestBucketAutoAllNilParams(t *testing.T) {
	res := builder.BucketAuto(nil, nil, nil, nil)
	require.Equal(t, operator.BucketAuto, res.GetKey())
	assert.Equal(t, bson.M{}, res.GetVal())
}

func TestCollStatsAllNilParams(t *testing.T) {
	res := builder.CollStats(nil, nil, nil)
	require.Equal(t, operator.CollStats, res.GetKey())
	assert.Equal(t, bson.M{}, res.GetVal())
}

func TestCurrentOpAllNilParams(t *testing.T) {
	res := builder.CurrentOp(nil, nil, nil, nil, nil)
	require.Equal(t, operator.CurrentOp, res.GetKey())
	assert.Equal(t, bson.M{}, res.GetVal())
}

func TestLookupAllNilParams(t *testing.T) {
	res := builder.Lookup(nil, nil, nil, nil)
	require.Equal(t, operator.Lookup, res.GetKey())
	assert.Equal(t, bson.M{}, res.GetVal())
}

func TestUncorrelatedLookupAllNilParams(t *testing.T) {
	res := builder.UncorrelatedLookup(nil, nil, nil, nil)
	require.Equal(t, operator.Lookup, res.GetKey())
	assert.Equal(t, bson.M{}, res.GetVal())
}

func TestMergeAllNilParams(t *testing.T) {
	res := builder.Merge(nil, nil, nil, nil, nil)
	require.Equal(t, operator.Merge, res.GetKey())
	assert.Equal(t, bson.M{}, res.GetVal())
}

func TestUnwindAllNilParams(t *testing.T) {
	res := builder.Unwind(nil, nil, nil)
	require.Equal(t, operator.Unwind, res.GetKey())
	assert.Equal(t, bson.M{}, res.GetVal())
}

func TestGroupWithNilParams(t *testing.T) {
	res := builder.Group(nil, nil)
	require.Equal(t, operator.Group, res.GetKey())
	// nil params map should result in just empty map (no _id since ID is nil)
	assert.Equal(t, bson.M{}, res.GetVal())
}

func TestGroupWithEmptyParams(t *testing.T) {
	res := builder.Group("$field", bson.M{})
	require.Equal(t, operator.Group, res.GetKey())
	assert.Equal(t, bson.M{"_id": "$field"}, res.GetVal())
}

func TestGroupWithNilValueInParams(t *testing.T) {
	res := builder.Group("$field", bson.M{"count": nil, "total": "$amount"})
	m := res.GetVal().(bson.M)
	// nil values should be filtered out by appendIfHasVal
	_, hasCount := m["count"]
	assert.False(t, hasCount, "nil value should be filtered out")
	assert.Equal(t, "$amount", m["total"])
	assert.Equal(t, "$field", m["_id"])
}

func TestBucketZeroValues(t *testing.T) {
	// Zero values (0, false, "") are NOT nil, so they should be included
	res := builder.Bucket(0, false, "", "output")
	m := res.GetVal().(bson.M)

	assert.Equal(t, 0, m["groupBy"])
	assert.Equal(t, false, m["boundaries"])
	assert.Equal(t, "", m["default"])
	assert.Equal(t, "output", m["output"])
}

func TestCurrentOpWithFalseValues(t *testing.T) {
	// false is not nil and should be included
	res := builder.CurrentOp(false, false, false, false, false)
	m := res.GetVal().(bson.M)

	assert.Equal(t, false, m["allUsers"])
	assert.Equal(t, false, m["idleConnections"])
	assert.Equal(t, false, m["idleCursors"])
	assert.Equal(t, false, m["idleSessions"])
	assert.Equal(t, false, m["localOps"])
}

func TestSWithNoOperators(t *testing.T) {
	result := builder.S()
	assert.Equal(t, bson.M{}, result)
}

func TestSWithMultipleOperators(t *testing.T) {
	op1 := builder.New("key1", "val1")
	op2 := builder.New("key2", "val2")
	result := builder.S(op1, op2)

	assert.Equal(t, bson.M{"key1": "val1", "key2": "val2"}, result)
}

func TestOperatorGetKeyAndGetVal(t *testing.T) {
	op := builder.New("testKey", bson.M{"a": 1})
	assert.Equal(t, "testKey", op.GetKey())
	assert.Equal(t, bson.M{"a": 1}, op.GetVal())
}

func TestSMapToMap(t *testing.T) {
	op1 := builder.New("$match", bson.M{"active": true})
	result := builder.S(op1)

	assert.Equal(t, bson.M{"$match": bson.M{"active": true}}, result)
}

func TestLookupPartialNil(t *testing.T) {
	res := builder.Lookup("users", nil, nil, "user_data")
	m := res.GetVal().(bson.M)

	assert.Equal(t, "users", m["from"])
	assert.Equal(t, "user_data", m["as"])
	_, hasLocal := m["localField"]
	_, hasForeign := m["foreignField"]
	assert.False(t, hasLocal, "nil localField should be omitted")
	assert.False(t, hasForeign, "nil foreignField should be omitted")
}

func TestMergePartialNil(t *testing.T) {
	res := builder.Merge("output", nil, nil, "replace", nil)
	m := res.GetVal().(bson.M)

	assert.Equal(t, "output", m["into"])
	assert.Equal(t, "replace", m["whenMatched"])
	_, hasOn := m["on"]
	_, hasLet := m["let"]
	_, hasWhenNot := m["whenNotMatched"]
	assert.False(t, hasOn)
	assert.False(t, hasLet)
	assert.False(t, hasWhenNot)
}

func TestUnwindPartialNil(t *testing.T) {
	res := builder.Unwind("$items", nil, true)
	m := res.GetVal().(bson.M)

	assert.Equal(t, "$items", m["path"])
	assert.Equal(t, true, m["preserveNullAndEmptyArrays"])
	_, hasIndex := m["includeArrayIndex"]
	assert.False(t, hasIndex, "nil includeArrayIndex should be omitted")
}

func TestReplaceRootWithNil(t *testing.T) {
	res := builder.ReplaceRoot(nil)
	m := res.GetVal().(bson.M)

	_, hasNewRoot := m["newRoot"]
	assert.False(t, hasNewRoot, "nil newRoot should be omitted")
}

func TestSampleWithZero(t *testing.T) {
	res := builder.Sample(0)
	m := res.GetVal().(bson.M)

	assert.Equal(t, 0, m["size"])
}
