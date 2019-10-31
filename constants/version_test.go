package constants

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Version(t *testing.T) {
	assert.Equal(t, "string", reflect.TypeOf(Version).String())
}
