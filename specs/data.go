package specs

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Data struct {
	Data *DataConfigurations
}

type DataConfigurations struct {
	Format       string    `yaml:"format"`
	SearchColumn []*Column `yaml:"search_column"`
	CreateColumn []*Column `yaml:"create_column"`
	TableColumn  []*Column `yaml:"table_column"`
	ExtraColumn  []*Column `yaml:"extra_column"`
	ButtonName   string    `yaml:"button_name"`
	Name         string    `yaml:"name"`
	Package      string    `yaml:"package"`
}

type Column struct {
	Prop       string `yaml:"prop"`
	Label      string `yaml:"label"`
	Type       string `yaml:"type"`
	ColumnName string `yaml:"column_name"`
}

func (data *Data) ToString() {
	fmt.Printf(data.Data.Format + "|")
	for _, item := range data.Data.SearchColumn {
		fmt.Println(item)

	}
	fmt.Printf(data.Data.ButtonName)

}

func loadDataFromDataDirectory(path string) (*Data, error) {
	data := newDefaultData()
	if content, err := ioutil.ReadFile(path + "/data.yaml"); err == nil {
		err = yaml.Unmarshal(content, data)
		if err != nil {
			return nil, err
		}
		return data, nil
	}

	return nil, errors.New("data file not found in " + path + ".")
}

func newDefaultData() *Data {

	return &Data{
		Data: &DataConfigurations{
			Format:       "mysql",
			SearchColumn: []*Column{},
			CreateColumn: []*Column{},
			TableColumn:  []*Column{},
			ExtraColumn:  []*Column{},
			ButtonName:   "",
			Name:         "",
		},
	}
}
