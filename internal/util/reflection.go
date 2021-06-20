package util

import "reflect"

// IsNil function determines whether the parameter is nil or not. If the value itself is not nil,
// reflection is used to determine whether the underlying value is nil.
// See https://play.golang.org/p/Isoo0CcAvr for an example.
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

// AnyNil returns true if any of the passed in parameters are nil, and returns false otherwise.
func AnyNil(values ...interface{}) bool {
	for _, val := range values {

		if IsNil(val) {
			return true
		}

	}

	return false
}
