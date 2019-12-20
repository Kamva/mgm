package builder

type Operator interface {
	GetKey() string
	GetVal() interface{}
}

type BaseOperator struct {
	key string
	val interface{}
}

func (operator *BaseOperator) GetKey() string {
	return operator.key
}

func (operator *BaseOperator) GetVal() interface{} {
	return operator.val
}

func New(key string, val interface{}) Operator {
	return &BaseOperator{
		key: key,
		val: val,
	}
}

// Ensure that BaseOperator implemented Operator
var _ Operator = &BaseOperator{}
