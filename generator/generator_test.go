package generator

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lwydyby/generator/specs"
	"github.com/lwydyby/generator/templates"
)

func TestGenerator(t *testing.T) {
	templatePath, err := filepath.Abs("../templates/example")
	assert.Nil(t, err)
	loadedTemplates, manifest, err := templates.LoadTemplates(templatePath)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(loadedTemplates))
	assert.Equal(t, "web", loadedTemplates["web"].ID)
	assert.Equal(t, "Go", manifest.Template.Format)

	fixtures, err := filepath.Abs("../templates/example")
	assert.Nil(t, err)
	specInfo, err := specs.LoadData(fixtures)
	assert.Nil(t, err)
	assert.NotNil(t, specInfo)



	codeGenerator := New()
	codeGenerator.Set(specInfo, loadedTemplates["web"])

	err = codeGenerator.Render()


	loadedTemplates["web"].Format = "Go"
	err = codeGenerator.Render()
	assert.Nil(t, err)
	//loadedTemplates["web"].UpdateOutputFilename("index.vue", "camel_case")
	//loadedTemplates["dao"].UpdateOutputFilename("user.java", "camel_case")
	loadedTemplates["web"].UpdateOutputFilePath("./test")
	loadedTemplates["dao"].UpdateOutputFilePath("./test")

	err = codeGenerator.Write()

	//err = os.RemoveAll("./test")
	assert.Nil(t, err)
	codeGenerator.Clear()
}
