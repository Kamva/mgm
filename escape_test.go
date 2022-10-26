package mgm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEscape(t *testing.T) {
	assert.Equal(t, "abc\uFF04def", Escape("abc$def"))
	assert.Equal(t, "abc\uFF0Edef", Escape("abc.def"))
	assert.Equal(t, "abc\uFF04def\uFF0Eghi", Escape("abc$def.ghi"))
}

func TestUnescape(t *testing.T) {
	assert.Equal(t, "abc$def", Unescape("abc\uFF04def"))
	assert.Equal(t, "abc.def", Unescape("abc\uFF0Edef"))
	assert.Equal(t, "abc$def.ghi", Unescape("abc\uFF04def\uFF0Eghi"))
}

func BenchmarkEscape(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Escape("abc$def")
	}
}

func BenchmarkUnescape(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Unescape("abc\uFF04def")
	}
}
