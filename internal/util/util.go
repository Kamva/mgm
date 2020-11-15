package util

// PanicErr panic passed error if it's not nil.
func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}
