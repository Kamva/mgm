package util

import (
	"reflect"
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func PanicErr(err error) {
	if err != nil && !reflect.ValueOf(err).IsNil() {
		panic(err)
	}
}

func InterfaceIsNil(val interface{}) bool {
	// We can not directly check interface is nil or
	// not, check this https://play.golang.org/p/Isoo0CcAvr.
	// Also we firstly check `val==nil` because reflection
	// can not get value of zero val.
	return val == nil || reflect.ValueOf(val).IsNil()
}

// AnyNil return true if exists any nil value in passed params.
func AnyNil(values ...interface{}) bool {
	for _, val := range values {

		if InterfaceIsNil(val) {
			return true
		}

	}

	return false
}
