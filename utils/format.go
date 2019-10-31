package utils

import (
	"regexp"
)

// ReplaceCurlyBracketWithSquare replaces CurlyBracket to SquareBracket.
func ReplaceCurlyBracketWithSquare(original string) string {
	converted := original
	converted = regexp.MustCompile(`\{ *`).ReplaceAllString(converted, "<")
	converted = regexp.MustCompile(` *}`).ReplaceAllString(converted, ">")

	return converted
}
