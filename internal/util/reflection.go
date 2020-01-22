package util

import "reflect"

// IsNil function check value is nil or no. To check real value of interface
// is nil or not, should using reflection, check this
// https://play.golang.org/p/Isoo0CcAvr. Firstly check
// `val==nil` because reflection can not get value of
// zero val.
func IsNil(val interface{}) (result bool) {

	if val == nil {
		return true
	}

	switch v := reflect.ValueOf(val); v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer,
		reflect.Interface, reflect.Slice:
		return v.IsNil()
	}

	return
}

// AnyNil return true if exists any nil value in passed params.
func AnyNil(values ...interface{}) bool {
	for _, val := range values {

		if IsNil(val) {
			return true
		}

	}

	return false
}
