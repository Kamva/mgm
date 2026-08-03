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

func TestEscapeEmptyString(t *testing.T) {
	assert.Equal(t, "", Escape(""))
}

func TestUnescapeEmptyString(t *testing.T) {
	assert.Equal(t, "", Unescape(""))
}

func TestEscapeNoSpecialChars(t *testing.T) {
	assert.Equal(t, "hello_world", Escape("hello_world"))
}

func TestUnescapeNoSpecialChars(t *testing.T) {
	assert.Equal(t, "hello_world", Unescape("hello_world"))
}

func TestEscapeOnlyDollar(t *testing.T) {
	assert.Equal(t, "\uFF04", Escape("$"))
}

func TestEscapeOnlyDot(t *testing.T) {
	assert.Equal(t, "\uFF0E", Escape("."))
}

func TestEscapeMultipleDollars(t *testing.T) {
	assert.Equal(t, "\uFF04\uFF04\uFF04", Escape("$$$"))
}

func TestEscapeMultipleDots(t *testing.T) {
	assert.Equal(t, "\uFF0E\uFF0E\uFF0E", Escape("..."))
}

func TestEscapeBothDollarAndDot(t *testing.T) {
	assert.Equal(t, "\uFF04\uFF0E", Escape("$."))
}

func TestEscapeUnicodeContent(t *testing.T) {
	// Non-dollar/dot unicode should pass through unchanged
	assert.Equal(t, "héllo", Escape("héllo"))
	assert.Equal(t, "日本語", Escape("日本語"))
}

func TestEscapeRoundTrip(t *testing.T) {
	originals := []string{
		"$set",
		"field.name",
		"$field.nested$value",
		"plain",
		"",
		"$",
		".",
		"a$b.c$d.e",
	}

	for _, original := range originals {
		escaped := Escape(original)
		unescaped := Unescape(escaped)
		assert.Equal(t, original, unescaped, "Round-trip failed for: %q", original)
	}
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
