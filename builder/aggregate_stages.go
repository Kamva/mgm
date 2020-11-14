package builder

import (
	f "github.com/kamva/mgm/v3/field"
	o "github.com/kamva/mgm/v3/operator"
	"go.mongodb.org/mongo-driver/bson"
)

// Bucket function return mongo $bucket operator to using in aggregates.
func Bucket(groupBy, boundaries, def, output interface{}) Operator {
	m := bson.M{}

	appendIfHasVal(m, f.GroupBy, groupBy)
	appendIfHasVal(m, f.Boundaries, boundaries)
	appendIfHasVal(m, f.Default, def)
	appendIfHasVal(m, f.Output, output)

	return New(o.Bucket, m)
}

// BucketAuto function return mongo $bucketAuto operator to using in aggregates.
func BucketAuto(groupBy, buckets, output, granularity interface{}) Operator {
	m := bson.M{}

	appendIfHasVal(m, f.GroupBy, groupBy)
	appendIfHasVal(m, f.Buckets, buckets)
	appendIfHasVal(m, f.Output, output)
	appendIfHasVal(m, f.Granularity, granularity)

	return New(o.BucketAuto, m)
}

// CollStats function return mongo $collStats operator to using in aggregates.
func CollStats(latencyStats, storageStats, count interface{}) Operator {
	m := bson.M{}

	appendIfHasVal(m, f.LatencyStats, latencyStats)
	appendIfHasVal(m, f.StorageStats, storageStats)
	appendIfHasVal(m, f.Count, count)

	return New(o.CollStats, m)
}

// CurrentOp function return mongo $currentOp operator to using in aggregates.
func CurrentOp(allUsers, idleConnections, idleCursors, idleSessions, localOps interface{}) Operator {
	m := bson.M{}

	appendIfHasVal(m, f.AllUsers, allUsers)
	appendIfHasVal(m, f.IdleConnections, idleConnections)
	appendIfHasVal(m, f.IdleCursors, idleCursors)
	appendIfHasVal(m, f.IdleSessions, idleSessions)
	appendIfHasVal(m, f.LocalOps, localOps)

	return New(o.CurrentOp, m)
}

// $geoNear,$graphLookup has many params, those functions
// will have too many params and do not make readable code.

// Group function return mongo $group operator to using in aggregates.
func Group(ID interface{}, params bson.M) Operator {
	m := bson.M{}

	appendIfHasVal(m, f.ID, ID)

	for key, val := range params {
		appendIfHasVal(m, key, val)
	}

	return New(o.Group, m)
}

// Lookup function return mongo $lookup operator to using in aggregates.
func Lookup(from, localField, foreignField, as interface{}) Operator {
	m := bson.M{}

	appendIfHasVal(m, f.From, from)
	appendIfHasVal(m, f.LocalField, localField)
	appendIfHasVal(m, f.ForeignField, foreignField)
	appendIfHasVal(m, f.As, as)

	return New(o.Lookup, m)
}

// UncorrelatedLookup function return mongo $lookup operator to using in aggregates.
func UncorrelatedLookup(from, let, pipeline, as interface{}) Operator {
	m := bson.M{}

	appendIfHasVal(m, f.From, from)
	appendIfHasVal(m, f.Let, let)
	appendIfHasVal(m, f.Pipeline, pipeline)
	appendIfHasVal(m, f.As, as)

	return New(o.Lookup, m)
}

// Merge function return mongo $merge operator to using in aggregates.
func Merge(into, on, let, whenMatched, whenNotMatched interface{}) Operator {
	m := bson.M{}

	appendIfHasVal(m, f.Into, into)
	appendIfHasVal(m, f.On, on)
	appendIfHasVal(m, f.Let, let)
	appendIfHasVal(m, f.WhenMatched, whenMatched)
	appendIfHasVal(m, f.WhenNotMatched, whenNotMatched)

	return New(o.Merge, m)
}

// ReplaceRoot function return mongo $replaceRoot operator to using in aggregates.
func ReplaceRoot(newRoot interface{}) Operator {
	m := bson.M{}

	appendIfHasVal(m, f.NewRoot, newRoot)

	return New(o.ReplaceRoot, m)
}

// Sample function return mongo sample operator to using in aggregates.
func Sample(size interface{}) Operator {
	m := bson.M{}

	appendIfHasVal(m, f.Size, size)

	return New(o.Sample, m)
}

// Unwind function return mongo $unwind operator to using in aggregates.
func Unwind(path, includeArrayIndex, preserveNullAndEmptyArrays interface{}) Operator {
	m := bson.M{}

	appendIfHasVal(m, f.Path, path)
	appendIfHasVal(m, f.IncludeArrayIndex, includeArrayIndex)
	appendIfHasVal(m, f.PreserveNullAndEmptyArrays, preserveNullAndEmptyArrays)

	return New(o.Unwind, m)
}
