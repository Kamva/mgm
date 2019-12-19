package operator

// Collection aggregation Stages
const (
	AddFields      = "$addFields"
	Bucket         = "$bucket"
	BucketAuto     = "$bucketAuto"
	CollStats      = "$collStats"
	Count          = "$count"
	Facet          = "$facet"
	GeoNear        = "$geoNear"
	GraphLookup    = "$graphLookup"
	Group          = "$group"
	IndexStats     = "$indexStats"
	Limit          = "$limit"
	ListSessions   = "$listSessions"
	Lookup         = "$lookup"
	Match          = "$match"
	Merge          = "$merge"
	Out            = "$out"
	PlanCacheStats = "$planCacheStats"
	Project        = "$project"
	Redact         = "$redact"
	ReplaceRoot    = "$replaceRoot"
	ReplaceWith    = "$replaceWith"
	Sample         = "$sample"
	// Set            = "$set" // Declared
	Skip = "$skip"
	// Sort           = "$sort" // Declared
	SortByCount = "$sortByCount"
	// Unset          = "$unset" // Declared
	Unwind = "$unwind"
)

// DB Aggregate stages
const (
	CurrentOp         = "$currentOp"
	ListLocalSessions = "$listLocalSessions"
)
