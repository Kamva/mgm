package util

import (
	"reflect"
)

// PanicErr panic passed error if it's not nil.
func PanicErr(err error) {
	if err != nil && !reflect.ValueOf(err).IsNil() {
		panic(err)
	}
}
