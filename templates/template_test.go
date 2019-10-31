package templates

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplate(t *testing.T) {
	fullPath, err := filepath.Abs("example")
	assert.Nil(t, err)

	templates, _, err := LoadTemplates(fullPath)
	assert.Nil(t, err)

	webTemplate := templates["web"]
	assert.NotNil(t, webTemplate)


	err = webTemplate.UpdateOutputFilePath("/test/path")
	assert.Nil(t, err)
	assert.NotNil(t, webTemplate)
	assert.Equal(t, webTemplate.OutputFilePath, "/test/path")
}
