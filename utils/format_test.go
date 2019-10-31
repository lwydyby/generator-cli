package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceCurlyBracketWithSquare(t *testing.T) {
	target := "/<bucketName>?stats"
	assert.Equal(t, target, ReplaceCurlyBracketWithSquare("/{bucketName}?stats"))
	assert.Equal(t, target, ReplaceCurlyBracketWithSquare("/{bucketName }?stats"))
	assert.Equal(t, target, ReplaceCurlyBracketWithSquare("/{ bucketName}?stats"))

	anotherTarget := "/<<bucketName>>?stats"
	assert.Equal(t, anotherTarget, ReplaceCurlyBracketWithSquare("/{{bucketName}}?stats"))
	assert.Equal(t, anotherTarget, ReplaceCurlyBracketWithSquare("/{{ bucketName }}?stats"))
}
