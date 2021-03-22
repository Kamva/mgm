package util

// PanicErr panics using the passed-in error if it's not nil.
func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}
