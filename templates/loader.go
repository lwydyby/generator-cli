package templates

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// LoadTemplates read the template manifest and load all templates info.
func LoadTemplates(templateDirectory string) (map[string]*Template, *ManifestConfigurations, error) {
	templates := map[string]*Template{}

	manifest, err := loadManifestFromTemplateDirectory(templateDirectory)
	if err != nil {
		return templates, nil, err
	}

	for templateID := range manifest.TemplateFiles {
		serviceTemplate, err := loadTemplate(templateDirectory, templateID, manifest)
		if err == nil {
			templates[templateID] = serviceTemplate
		}
	}

	return templates, manifest, nil
}

func loadTemplate(templateDirectory, templateID string, manifest *ManifestConfigurations) (*Template, error) {
	templateConfiguration := manifest.TemplateFiles[templateID]
	if templateConfiguration == nil {
		return nil, fmt.Errorf(`configuration of template "%s" not found`, templateID)
	}

	templateFilePath := templateDirectory + "/" + templateConfiguration.FilePath
	templateFileContent, err := ioutil.ReadFile(templateFilePath)
	if err != nil {
		return nil, err
	}

	template := &Template{
		Filename:       filepath.Base(templateFilePath),
		FileDirectory:  filepath.Dir(templateConfiguration.FilePath),
		FilePath:       templateFilePath,
		Format:         manifest.Template.Format,
		OutputFilename: manifest.TemplateFiles[templateID].OutputFileNaming,
		FileContent:    string(templateFileContent),
		IsNeedGenerate: true,
		ID:             templateID,
	}

	return template, nil
}
