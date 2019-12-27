package builder

// Operator is interface that should implement by each struct
// that want to be an operator.
type Operator interface {
	GetKey() string
	GetVal() interface{}
}

// BaseOperator is simple base operator that implemented Operator.
type BaseOperator struct {
	key string
	val interface{}
}

// GetKey function return operator's key.
func (operator *BaseOperator) GetKey() string {
	return operator.key
}

// GetVal method return operator's value.
func (operator *BaseOperator) GetVal() interface{} {
	return operator.val
}

// New function return new Operator
func New(key string, val interface{}) Operator {
	return &BaseOperator{
		key: key,
		val: val,
	}
}

// Ensure that BaseOperator implemented Operator
var _ Operator = &BaseOperator{}
