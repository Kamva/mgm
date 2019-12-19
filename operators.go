package mgm

type Operator interface {
	GetKey() string
	GetVal() string
}

// TODO: define operators.