package builder

import (
	f "github.com/Kamva/mgm/field"
	o "github.com/Kamva/mgm/operator"
	"go.mongodb.org/mongo-driver/bson"
)

func Bucket(groupBy, boundaries, def, output interface{}) Operator {
	m := bson.M{}

	appendIfHasVal(m, f.GroupBy, groupBy)
	appendIfHasVal(m, f.Boundaries, boundaries)
	appendIfHasVal(m, f.Default, def)
	appendIfHasVal(m, f.Output, output)

	return New(o.Bucket, m)
}

func BucketAuto(groupBy, buckets, output, granularity interface{}) Operator {
	m := bson.M{}

	appendIfHasVal(m, f.GroupBy, groupBy)
	appendIfHasVal(m, f.Buckets, buckets)
	appendIfHasVal(m, f.Output, output)
	appendIfHasVal(m, f.Granularity, granularity)

	return New(o.BucketAuto, m)
}

func CollStats(latencyStats, storageStats, count interface{}) Operator {
	m := bson.M{}

	appendIfHasVal(m, f.LatencyStats, latencyStats)
	appendIfHasVal(m, f.StorageStats, storageStats)
	appendIfHasVal(m, f.Count, count)

	return New(o.CollStats, m)
}

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

func Group(Id interface{}, params bson.M) Operator {
	m := bson.M{}

	appendIfHasVal(m, f.Id, Id)

	for key, val := range params {
		appendIfHasVal(m, key, val)
	}

	return New(o.Group, m)
}

func Lookup(from, localField, foreignField, as interface{}) Operator {
	m := bson.M{}

	appendIfHasVal(m, f.From, from)
	appendIfHasVal(m, f.LocalField, localField)
	appendIfHasVal(m, f.ForeignField, foreignField)
	appendIfHasVal(m, f.As, as)

	return New(o.Lookup, m)
}

func UncorrelatedLookup(from, let, pipeline, as interface{}) Operator {
	m := bson.M{}

	appendIfHasVal(m, f.From, from)
	appendIfHasVal(m, f.Let, let)
	appendIfHasVal(m, f.Pipeline, pipeline)
	appendIfHasVal(m, f.As, as)

	return New(o.Lookup, m)
}


func Merge(into, on, let, whenMatched, whenNotMatched interface{}) Operator {
	m := bson.M{}

	appendIfHasVal(m, f.Into, into)
	appendIfHasVal(m, f.On, on)
	appendIfHasVal(m, f.Let, let)
	appendIfHasVal(m, f.WhenMatched, whenMatched)
	appendIfHasVal(m, f.WhenNotMatched, whenNotMatched)

	return New(o.Merge, m)
}

func ReplaceRoot(newRoot interface{}) Operator {
	m := bson.M{}

	appendIfHasVal(m, f.NewRoot, newRoot)

	return New(o.ReplaceRoot, m)
}

func Sample(size interface{}) Operator {
	m := bson.M{}

	appendIfHasVal(m, f.Size, size)

	return New(o.Sample, m)
}

func Unwind(path, includeArrayIndex, preserveNullAndEmptyArrays interface{}) Operator {
	m := bson.M{}

	appendIfHasVal(m, f.Path, path)
	appendIfHasVal(m, f.IncludeArrayIndex, includeArrayIndex)
	appendIfHasVal(m, f.PreserveNullAndEmptyArrays, preserveNullAndEmptyArrays)

	return New(o.Unwind, m)
}
