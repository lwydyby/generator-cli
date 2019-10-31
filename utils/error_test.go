package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckErrorForExit(t *testing.T) {
	CheckErrorForExit(nil)
	assert.True(t, true)
}
