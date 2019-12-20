package builder_test

import (
	"github.com/Kamva/mgm/builder"
	"github.com/Kamva/mgm/field"
	"github.com/Kamva/mgm/operator"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

type TestLookupData struct {
	inputs []interface{}
	result interface{}
	hasErr bool
}

func TestGetNewOperator(t *testing.T) {
	res := builder.New("foo", "bar")
	require.Equal(t, res.GetKey(), "foo")
	require.Equal(t, res.GetVal(), "bar")
}

func TestBucket(t *testing.T) {
	dataItems := []TestLookupData{
		{
			inputs: []interface{}{"foo", "bar", "baz", "qux"},
			result: builder.New(operator.Bucket, bson.M{
				field.GroupBy:    "foo",
				field.Boundaries: "bar",
				field.Default:    "baz",
				field.Output:     "qux",
			}),
		},
	}

	for _, item := range dataItems {
		res := builder.Bucket(item.inputs[0], item.inputs[1], item.inputs[2], item.inputs[3])
		require.Equal(t, item.result, res)
	}
}
func TestBucketAuto(t *testing.T) {
	dataItems := []TestLookupData{
		{
			inputs: []interface{}{"foo", "bar", "baz", "qux"},
			result: builder.New(operator.BucketAuto, bson.M{
				field.GroupBy:     "foo",
				field.Buckets:     "bar",
				field.Output:      "baz",
				field.Granularity: "qux",
			}),
		},
		{
			inputs: []interface{}{"foo", 5, nil, "qux"},
			result: builder.New(operator.BucketAuto, bson.M{
				field.GroupBy:     "foo",
				field.Buckets:     5,
				field.Granularity: "qux",
			}),
		},
	}

	for _, item := range dataItems {
		res := builder.BucketAuto(item.inputs[0], item.inputs[1], item.inputs[2], item.inputs[3])
		require.Equal(t, item.result, res)
	}
}

func TestCollStats(t *testing.T) {
	dataItems := []TestLookupData{
		{
			inputs: []interface{}{"foo", "bar", "baz"},
			result: builder.New(operator.CollStats, bson.M{
				field.LatencyStats: "foo",
				field.StorageStats: "bar",
				field.Count:        "baz",
			}),
		},
		{
			inputs: []interface{}{nil, nil, bson.M{}},
			result: builder.New(operator.CollStats, bson.M{
				field.Count: bson.M{},
			}),
		},
	}

	for _, item := range dataItems {
		res := builder.CollStats(item.inputs[0], item.inputs[1], item.inputs[2])
		require.Equal(t, item.result, res)
	}
}

func TestCurrentOp(t *testing.T) {
	dataItems := []TestLookupData{
		{
			inputs: []interface{}{"foo", "bar", "baz", "qux", "quux"},
			result: builder.New(operator.CurrentOp, bson.M{
				field.AllUsers:        "foo",
				field.IdleConnections: "bar",
				field.IdleCursors:     "baz",
				field.IdleSessions:    "qux",
				field.LocalOps:        "quux",
			}),
		},
		{
			inputs: []interface{}{true, nil, nil, true, nil},
			result: builder.New(operator.CurrentOp, bson.M{
				field.AllUsers:     true,
				field.IdleSessions: true,
			}),
		},
	}

	for _, item := range dataItems {
		res := builder.CurrentOp(item.inputs[0], item.inputs[1], item.inputs[2],
			item.inputs[3], item.inputs[4])

		require.Equal(t, item.result, res)
	}
}
func TestGroup(t *testing.T) {
	dataItems := []TestLookupData{
		{
			inputs: []interface{}{"foo", bson.M{"bar": "baz", "qux": "quux"}},
			result: builder.New(operator.Group, bson.M{
				field.Id: "foo",
				"bar":    "baz",
				"qux":    "quux",
			}),
		},

		{
			inputs: []interface{}{"foo", bson.M{}},
			result: builder.New(operator.Group, bson.M{
				field.Id: "foo",
			}),
		},
	}

	for _, item := range dataItems {
		res := builder.Group(item.inputs[0], item.inputs[1].(bson.M))
		require.Equal(t, item.result, res)
	}
}

func TestLookup(t *testing.T) {
	dataItems := []TestLookupData{
		{
			inputs: []interface{}{"foo", "bar", "baz", "qux"},
			result: builder.New(operator.Lookup, bson.M{
				field.From:         "foo",
				field.LocalField:   "bar",
				field.ForeignField: "baz",
				field.As:           "qux",
			}),
		},
		{
			inputs: []interface{}{"foo", "bar", "baz", nil},
			result: builder.New(operator.Lookup, bson.M{
				field.From:         "foo",
				field.LocalField:   "bar",
				field.ForeignField: "baz",
			}),
		},
	}

	for _, item := range dataItems {
		res := builder.Lookup(item.inputs[0], item.inputs[1], item.inputs[2], item.inputs[3])
		require.Equal(t, item.result, res)
	}
}

func TestUncorrelatedLookup(t *testing.T) {
	dataItems := []TestLookupData{
		{
			inputs: []interface{}{"foo", "bar", "baz", "qux"},
			result: builder.New(operator.Lookup, bson.M{
				field.From:     "foo",
				field.Let:      "bar",
				field.Pipeline: "baz",
				field.As:       "qux",
			}),
		},
		{
			inputs: []interface{}{"foo", "bar", "baz", nil},
			result: builder.New(operator.Lookup, bson.M{
				field.From:     "foo",
				field.Let:      "bar",
				field.Pipeline: "baz",
			}),
		},
	}

	for _, item := range dataItems {
		res := builder.UncorrelatedLookup(item.inputs[0], item.inputs[1], item.inputs[2], item.inputs[3])
		require.Equal(t, item.result, res)
	}
}

func TestMerge(t *testing.T) {
	dataItems := []TestLookupData{
		{
			inputs: []interface{}{"foo", "bar", "baz", "qux", "quux"},
			result: builder.New(operator.Merge, bson.M{
				field.Into:           "foo",
				field.On:             "bar",
				field.Let:            "baz",
				field.WhenMatched:    "qux",
				field.WhenNotMatched: "quux",
			}),
		},
		{
			inputs: []interface{}{"myOutput", "_id", nil, "replace", "insert"},
			result: builder.New(operator.Merge, bson.M{
				field.Into:           "myOutput",
				field.On:             "_id",
				field.WhenMatched:    "replace",
				field.WhenNotMatched: "insert",
			}),
		},
	}

	for _, item := range dataItems {
		res := builder.Merge(item.inputs[0], item.inputs[1], item.inputs[2], item.inputs[3], item.inputs[4])
		require.Equal(t, item.result, res)
	}
}

func TestReplaceRoot(t *testing.T) {
	dataItems := []TestLookupData{
		{
			inputs: []interface{}{"foo"},
			result: builder.New(operator.ReplaceRoot, bson.M{
				field.NewRoot: "foo",
			}),
		},
		{
			inputs: []interface{}{"$name"},
			result: builder.New(operator.ReplaceRoot, bson.M{
				field.NewRoot: "$name",
			}),
		},
	}

	for _, item := range dataItems {
		res := builder.ReplaceRoot(item.inputs[0])
		require.Equal(t, item.result, res)
	}
}

func TestSample(t *testing.T) {
	dataItems := []TestLookupData{
		{
			inputs: []interface{}{45},
			result: builder.New(operator.Sample, bson.M{
				field.Size: 45,
			}),
		},
	}

	for _, item := range dataItems {
		res := builder.Sample(item.inputs[0])
		require.Equal(t, item.result, res)
	}
}

func TestUnwind(t *testing.T) {
	dataItems := []TestLookupData{
		{
			inputs: []interface{}{"foo", "bar", "baz"},
			result: builder.New(operator.Unwind, bson.M{
				field.Path:                       "foo",
				field.IncludeArrayIndex:          "bar",
				field.PreserveNullAndEmptyArrays: "baz",
			}),
		},
		{
			inputs: []interface{}{"$sizes", nil, nil},
			result: builder.New(operator.Unwind, bson.M{
				field.Path: "$sizes",
			}),
		},
	}

	for _, item := range dataItems {
		res := builder.Unwind(item.inputs[0], item.inputs[1], item.inputs[2])
		require.Equal(t, item.result, res)
	}
}

func TestGetSimpleMap(t *testing.T) {
	dataItems := []TestLookupData{
		{
			inputs: []interface{}{"foo", "bar", "baz", "qux"},
			result: builder.S{
				builder.New(operator.Lookup, bson.M{
					field.From:     "foo",
					field.Let:      "bar",
					field.Pipeline: "baz",
					field.As:       "qux",
				}),
			},
		},
		{
			inputs: []interface{}{"foo", "bar", "baz", nil},
			result: builder.S{
				builder.New(operator.Lookup, bson.M{
					field.From:     "foo",
					field.Let:      "bar",
					field.Pipeline: "baz",
				}),
			},
		},
	}

	for _, item := range dataItems {
		res := builder.S{
			builder.UncorrelatedLookup(item.inputs[0], item.inputs[1], item.inputs[2], item.inputs[3]),
		}

		require.Equal(t, item.result, res)
	}
}
