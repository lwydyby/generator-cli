package generator

import (
	"bytes"
	"fmt"
	"github.com/lwydyby/generator/specs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/lwydyby/generator/templates"
)

// Generator will load a API spec file and covert it to corresponding programing code.
type Generator struct {
	data  *specs.Data
	template *templates.Template

	code string
}

// New creates a generator.
func New() *Generator {
	return &Generator{}
}

// Set set properties of a generator.
func (g *Generator) Set(data *specs.Data, template *templates.Template) {
	g.data = data
	g.template = template
	g.code = ""
}

// Clear clear the generator's properties.
func (g *Generator) Clear() {
	g.data = nil
	g.template = nil
	g.code = ""
}

// Run will render and write code.
func (g *Generator) Run() error {
	err := g.Render()
	if err != nil {
		return err
	}

	err = g.Write()
	if err != nil {
		return err
	}

	return nil
}

// Render coverts API spec data content to programing code.
func (g *Generator) Render() error {
	if g.template.IsNeedGenerate {
		switch g.template.Format {
		case "Go":
			buffer := bytes.Buffer{}
			target := template.Must(template.New("Template").Funcs(funcMap).Parse(g.template.FileContent))

			var err error
			err = target.Execute(&buffer, g.data)
			if err != nil {
				return err
			}

			g.code = buffer.String()
		default:
			return fmt.Errorf("Template format not supported: \"%s\"", g.template.Format)
		}
	} else {
		g.code = string(g.template.FileContent)
	}

	return nil
}

// Write writes the generated code into the target file.
func (g *Generator) Write() error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	fmt.Println("Generating to: ." + strings.Replace(g.template.OutputFilePath, pwd, "", -1))

	err = os.MkdirAll(filepath.Dir(g.template.OutputFilePath), 0755)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(g.template.OutputFilePath, []byte(g.code), 0644)
	if err != nil {
		return err
	}

	return nil
}
