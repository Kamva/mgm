package operator

// Arithmetic Expression Operators
const (
	Abs    = "$abs"
	Add    = "$add"
	Ceil   = "$ceil"
	Divide = "$divide"
	Exp    = "$exp"
	Floor  = "$floor"
	Ln     = "$ln"
	Log    = "$log"
	Log10  = "$log10"
	// Mod      = "$mod" // Declared
	Multiply = "$multiply"
	Pow      = "$pow"
	Round    = "$round"
	Sqrt     = "$sqrt"
	Subtract = "$subtract"
	Trunc    = "$trunc"
)

// Array Expression Operators
const (
	ArrayToObject = "$arrayToObject"
	ConcatArrays  = "$concatArrays"
	Filter        = "$filter"
	// In            = "$in" // Declared
	IndexOfArray  = "$indexOfArray"
	IsArray       = "$isArray"
	Map           = "$map"
	ObjectToArray = "$objectToArray"
	Range         = "$range"
	Reduce        = "$reduce"
	ReverseArray  = "$reverseArray"
	// Size          = "$size" // Declared
	// Slice         = "$slice" // Declared
	Zip = "$zip"
)

// Boolean Expression Operators
const (
// And = "$and" // Declared
// Not = "$not" // Declared
// Or  = "$or" // Declared
)

// Comparison Expression Operators
const (
	Cmp = "$cmp"
	//Eq  = "$eq" // Declared
	//Gt  = "$gt" // Declared
	//Gte = "$gte" // Declared
	//Lt  = "$lt" // Declared
	//Lte = "$lte" // Declared
	//Ne  = "$ne" // Declared
)

// Conditional Expression Operators
const (
	Cond   = "$cond"
	IfNull = "$ifNull"
	Switch = "$switch"
)

// Date Expression Operators
const (
	DateFromParts  = "$dateFromParts"
	DateFromString = "$dateFromString"
	DateToParts    = "$dateToParts"
	DateToString   = "$dateToString"
	DayOfMonth     = "$dayOfMonth"
	DayOfWeek      = "$dayOfWeek"
	DayOfYear      = "$dayOfYear"
	Hour           = "$hour"
	IsoDayOfWeek   = "$isoDayOfWeek"
	IsoWeek        = "$isoWeek"
	IsoWeekYear    = "$isoWeekYear"
	Millisecond    = "$millisecond"
	Minute         = "$minute"
	Month          = "$month"
	Second         = "$second"
	ToDate         = "$toDate"
	Week           = "$week"
	Year           = "$year"
)

// Literal Expression Operator
const (
	Literal = "$literal"
)

// Object Expression Operators
const (
	MergeObjects = "$mergeObjects"
	// ObjectToArray = "$objectToArray" // Declared
)

// Set Expression Operators
const (
	AllElementsTrue = "$allElementsTrue"
	AnyElementTrue  = "$anyElementTrue"
	SetDifference   = "$setDifference"
	SetEquals       = "$setEquals"
	SetIntersection = "$setIntersection"
	SetIsSubset     = "$setIsSubset"
	SetUnion        = "$setUnion"
)

// String Expression Operators
const (
	Concat = "$concat"
	// DateFromString = "$dateFromString" // Declared
	// DateToString   = "$dateToString" // Declared
	IndexOfBytes = "$indexOfBytes"
	IndexOfCP    = "$indexOfCP"
	Ltrim        = "$ltrim"
	RegexFind    = "$regexFind"
	RegexFindAll = "$regexFindAll"
	RegexMatch   = "$regexMatch"
	Rtrim        = "$rtrim"
	Split        = "$split"
	StrLenBytes  = "$strLenBytes"
	StrLenCP     = "$strLenCP"
	Strcasecmp   = "$strcasecmp"
	Substr       = "$substr"
	SubstrBytes  = "$substrBytes"
	SubstrCP     = "$substrCP"
	ToLower      = "$toLower"
	ToString     = "$toString"
	Trim         = "$trim"
	ToUpper      = "$toUpper"
)

// Text Expression Operator
const (
// Meta = "$meta" // Declared
)

// Trigonometry Expression Operators
const (
	Sin              = "$sin"
	Cos              = "$cos"
	Tan              = "$tan"
	Asin             = "$asin"
	Acos             = "$acos"
	Atan             = "$atan"
	Atan2            = "$atan2"
	Asinh            = "$asinh"
	Acosh            = "$acosh"
	Atanh            = "$atanh"
	DegreesToRadians = "$degreesToRadians"
	RadiansToDegrees = "$radiansToDegrees"
)

// Type Expression Operators
const (
	Convert = "$convert"
	ToBool  = "$toBool"
	//ToDate     = "$toDate" // Declared
	ToDecimal  = "$toDecimal"
	ToDouble   = "$toDouble"
	ToInt      = "$toInt"
	ToLong     = "$toLong"
	ToObjectID = "$toObjectId"
	//ToString   = "$toString" // Declared
	//Type       = "$type" // Declared
)

// Accumulators ($group)
const (
	// AddToSet     = "$addToSet" // Declared
	Avg   = "$avg"
	First = "$first"
	Last  = "$last"
	// Max          = "$max" // Declared
	// MergeObjects = "$mergeObjects" // Declared
	// Min          = "$min" // Declared
	// Push         = "$push" // Declared
	StdDevPop  = "$stdDevPop"
	StdDevSamp = "$stdDevSamp"
	Sum        = "$sum"
)

// Accumulators (in Other Stages)
const (
// Avg        = "$avg" // Declared
// Max        = "$max" // Declared
// Min        = "$min" // Declared
// StdDevPop  = "$stdDevPop" // Declared
// StdDevSamp = "$stdDevSamp" // Declared
// Sum        = "$sum" // Declared
)

// Variable Expression Operators
const (
	Let = "$let"
)
