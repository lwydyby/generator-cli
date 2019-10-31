package specs


// LoadTemplates read the template manifest and load all templates info.
func LoadData(codeSpecFile string) (*Data,error) {
	data, err := loadDataFromDataDirectory(codeSpecFile)
	return data, err
}