package operator

// Fields
const (
	CurrentDate = "$currentDate"
	Inc         = "$inc"
	Min         = "$min"
	Max         = "$max"
	Mul         = "$mul"
	Rename      = "$rename"
	Set         = "$set"
	SetOnInsert = "$setOnInsert"
	Unset       = "$unset"
)

// Array Operators
const (
	// $: Act as a modifier
	// $[]: Act as a modifier
	// $[<identifier>]: Act as a modifier

	AddToSet = "$addToSet"
	Pop      = "$pop"
	Pull     = "$pull"
	Push     = "$push"
	PullAll  = "$pullAll"
)

// Array modifiers
const (
	Each     = "$each"
	Position = "$position"
	// Slice    = "$slice" // Declared
	Sort = "$sort"
)

// Array bitwise
const (
	Bit = "$bit"
)
