package util

import (
	"reflect"
)


func PanicErr(err error) {
	if err != nil && !reflect.ValueOf(err).IsNil() {
		panic(err)
	}
}

