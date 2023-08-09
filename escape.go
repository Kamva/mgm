package mgm

import (
	"strings"
)

var escape = strings.NewReplacer("$", "\uFF04", ".", "\uFF0E")
var unescape = strings.NewReplacer("\uFF04", "$", "\uFF0E", ".")

// Escape escapes the mongo key for . and $ characters.
func Escape(key string) string {
	return escape.Replace(key)
}

// Unescape unescapes the mongo key for . and $ characters.
func Unescape(key string) string {
	return unescape.Replace(key)
}
