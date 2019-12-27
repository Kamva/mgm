package operator

// Comparison
const (
	Eq  = "$eq"
	Gt  = "$gt"
	Gte = "$gte"
	In  = "$in"
	Lt  = "$lt"
	Lte = "$lte"
	Ne  = "$ne"
	Nin = "$nin"
)

// Logical
const (
	And = "$and"
	Not = "$not"
	Nor = "$nor"
	Or  = "$or"
)

// Element
const (
	Exists = "$exists"
	Type   = "$type"
)

// Evaluation
const (
	Expr       = "$expr"
	JSONSchema = "$jsonSchema"
	Mod        = "$mod"
	Regex      = "$regex"
	Text       = "$text"
	Where      = "$where"
)

// Geo spatial
const (
	GeoIntersects = "$geoIntersects"
	GeoWithin     = "$geoWithin"
	Near          = "$near"
	NearSphere    = "$nearSphere"
)

// Array
const (
	All       = "$all"
	ElemMatch = "$elemMatch"
	Size      = "$size"
)

// Bitwise
const (
	BitsAllClear = "$bitsAllClear"
	BitsAllSet   = "$bitsAllSet"
	BitsAnyClear = "$bitsAnyClear"
	BitsAnySet   = "$bitsAnySet"
)

// Comments
const (
	Comment = "$comment"
)

// Projection operators
const (
	Dollar = "$"
	// ElemMatch = "$elemMatch" // Declared
	Meta  = "$meta"
	Slice = "$slice"
)
