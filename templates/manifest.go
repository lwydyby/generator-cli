package templates

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

// ManifestConfigurations holds the information of manifest file.
type ManifestConfigurations struct {
	Template        *templateConfigurations                `json:"template" yaml:"template"`
	TemplateFiles   map[string]*templateFileConfigurations `json:"template_files" yaml:"template_files"`

}

type templateConfigurations struct {
	Format string `json:"format" yaml:"format"`
}



type templateFileConfigurations struct {
	FilePath         string                   `json:"file_path" yaml:"file_path"`
	OutputFileNaming string `json:"output_file_naming" yaml:"output_file_naming"`
}

func loadManifestFromTemplateDirectory(path string) (*ManifestConfigurations, error) {
	manifest := newDefaultManifest()

	if content, err := ioutil.ReadFile(path + "/setting.json"); err == nil {
		err = json.Unmarshal(content, manifest)
		if err != nil {
			return nil, err
		}
		return manifest, nil
	}

	if content, err := ioutil.ReadFile(path + "/setting.yaml"); err == nil {
		err = yaml.Unmarshal(content, manifest)
		if err != nil {
			return nil, err
		}
		return manifest, nil
	}

	return nil, errors.New("Template manifest file not found in " + path + ".")
}

func newDefaultManifest() *ManifestConfigurations {
	return &ManifestConfigurations{
		Template: &templateConfigurations{
			Format: "Go",
		},
		TemplateFiles: map[string]*templateFileConfigurations{
			"web": {
				FilePath:         "web.tmpl",
				OutputFileNaming: "",
			},
			"dao": {
				FilePath:         "dao.tmpl",
				OutputFileNaming: "",
			},
			"api": {
				FilePath:         "api.tmpl",
				OutputFileNaming: "",
			},
			"service": {
				FilePath:         "service.tmpl",
				OutputFileNaming: "",
			},
			"controller": {
				FilePath:         "controller.tmpl",
				OutputFileNaming: "",
			},
			"entity": {
				FilePath:         "entity.tmpl",
				OutputFileNaming: "",
			},
			"search": {
				FilePath:         "search.tmpl",
				OutputFileNaming: "",
			},


		},
	}
}
