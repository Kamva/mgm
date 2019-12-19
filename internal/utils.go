package internal

import (
	"reflect"
)

func PanicErr(err error) {
	if err != nil && !reflect.ValueOf(err).IsNil() {
		panic(err)
	}
}

// AnyNil return true if exists any nil value in passed params.
func AnyNil(values ...interface{}) bool {
	for _, val := range values {

		// We can not directly check interface is nil or
		// not, check this https://play.golang.org/p/Isoo0CcAvr.
		// Also we firstly check `val==nil` because reflection
		// can not get value of zero val.
		if val == nil || reflect.ValueOf(val).IsNil() {
			return true
		}

	}

	return false
}
