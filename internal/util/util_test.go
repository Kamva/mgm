package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNilWithNil(t *testing.T) {
	assert.True(t, IsNil(nil))
}

func TestIsNilWithNonNil(t *testing.T) {
	assert.False(t, IsNil("hello"))
	assert.False(t, IsNil(42))
	assert.False(t, IsNil(true))
}

func TestIsNilWithNilPointer(t *testing.T) {
	var p *int
	assert.True(t, IsNil(p))
}

func TestIsNilWithNonNilPointer(t *testing.T) {
	v := 42
	assert.False(t, IsNil(&v))
}

func TestIsNilWithNilSlice(t *testing.T) {
	var s []int
	assert.True(t, IsNil(s))
}

func TestIsNilWithEmptySlice(t *testing.T) {
	s := []int{}
	assert.False(t, IsNil(s))
}

func TestIsNilWithNilMap(t *testing.T) {
	var m map[string]int
	assert.True(t, IsNil(m))
}

func TestIsNilWithEmptyMap(t *testing.T) {
	m := map[string]int{}
	assert.False(t, IsNil(m))
}

func TestIsNilWithNilFunc(t *testing.T) {
	var f func()
	assert.True(t, IsNil(f))
}

func TestIsNilWithFunc(t *testing.T) {
	f := func() {}
	assert.False(t, IsNil(f))
}

func TestIsNilWithNilChan(t *testing.T) {
	var c chan int
	assert.True(t, IsNil(c))
}

func TestIsNilWithChan(t *testing.T) {
	c := make(chan int)
	assert.False(t, IsNil(c))
}

func TestIsNilWithNilInterface(t *testing.T) {
	var i interface{}
	assert.True(t, IsNil(i))
}

func TestIsNilWithZeroValues(t *testing.T) {
	assert.False(t, IsNil(0))
	assert.False(t, IsNil(false))
	assert.False(t, IsNil(""))
}

func TestAnyNilAllNil(t *testing.T) {
	assert.True(t, AnyNil(nil, nil, nil))
}

func TestAnyNilNoneNil(t *testing.T) {
	assert.False(t, AnyNil("a", 1, true))
}

func TestAnyNilMixed(t *testing.T) {
	assert.True(t, AnyNil("a", nil, true))
}

func TestAnyNilEmpty(t *testing.T) {
	assert.False(t, AnyNil())
}

func TestAnyNilWithNilPointer(t *testing.T) {
	var p *int
	assert.True(t, AnyNil("hello", p, 42))
}

func TestToSnakeCaseSimple(t *testing.T) {
	assert.Equal(t, "blog_post", ToSnakeCase("BlogPost"))
}

func TestToSnakeCaseSingleWord(t *testing.T) {
	assert.Equal(t, "book", ToSnakeCase("Book"))
}

func TestToSnakeCaseAlreadySnake(t *testing.T) {
	assert.Equal(t, "blog_post", ToSnakeCase("blog_post"))
}

func TestToSnakeCaseWithAcronym(t *testing.T) {
	assert.Equal(t, "html_parser", ToSnakeCase("HTMLParser"))
}

func TestToSnakeCaseMultipleWords(t *testing.T) {
	assert.Equal(t, "my_long_struct_name", ToSnakeCase("MyLongStructName"))
}

func TestToSnakeCaseEmpty(t *testing.T) {
	assert.Equal(t, "", ToSnakeCase(""))
}

func TestToSnakeCaseLowercase(t *testing.T) {
	assert.Equal(t, "lowercase", ToSnakeCase("lowercase"))
}

func TestToSnakeCaseAllCaps(t *testing.T) {
	assert.Equal(t, "api", ToSnakeCase("API"))
}

func TestPanicErrWithNil(t *testing.T) {
	assert.NotPanics(t, func() {
		PanicErr(nil)
	})
}

func TestPanicErrWithError(t *testing.T) {
	assert.Panics(t, func() {
		PanicErr(assert.AnError)
	})
}
