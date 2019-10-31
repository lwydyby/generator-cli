package cmds

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lwydyby/generator/specs"

	"github.com/lwydyby/generator/templates"

	"github.com/lwydyby/generator/constants"
	"github.com/lwydyby/generator/generator"
	"github.com/lwydyby/generator/utils"
)

var (
	//版本
	flagVersion bool
	//要生成页面数据文件地址
	codeSpecFile string
	//模板所在文件夹
	codeTemplateDirectory string
	//输出文件夹
	codeOutputDirectory string
)

func init() {
	RootCMD.Flags().BoolVarP(
		&flagVersion, "version", "v", false,
		"Show version.",
	)
	RootCMD.Flags().StringVarP(
		&codeSpecFile, "file", "f", "",
		"Specify the specs file.",
	)
	RootCMD.Flags().StringVarP(
		&codeTemplateDirectory, "template", "t", "",
		"Specify template files directory.",
	)
	RootCMD.Flags().StringVarP(
		&codeOutputDirectory, "output", "o", "",
		"Specify the output directory.",
	)
}

// RootCMD represents the base command when called without any subcommands.
var RootCMD = &cobra.Command{
	Use:   "generator cli",
	Short: "A code generator for spring boot program.",
	Long: `A code generator for spring boot program.
For example:
  $ generator -f ./data.json
          -t ./templates \
          -o ./output
`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if flagVersion {
			return nil
		}

		if codeSpecFile == "" {
			return errors.New("please specify the data file")
		}
		if _, err := os.Stat(codeSpecFile); err != nil {
			return errors.New("please make sure the data file exists")
		}

		if codeTemplateDirectory == "" {
			return errors.New("please specify templates directory")
		}
		if _, err := os.Stat(codeTemplateDirectory); err != nil {
			return errors.New("please make sure the templates directory exists")
		}

		if codeOutputDirectory == "" {
			return errors.New("please specify output files directory")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if flagVersion {
			fmt.Println("generator version " + constants.Version)
			return
		}

		loadedTemplates, _, err := templates.LoadTemplates(codeTemplateDirectory)
		utils.CheckErrorForExit(err)
		fmt.Println("Loaded templates from " + codeTemplateDirectory)
		fmt.Println(len(loadedTemplates), "template(s) detected.")

		spec, err := specs.LoadData(codeSpecFile)
		utils.CheckErrorForExit(err)
		fmt.Printf("Loaded data file %s \n\n", codeSpecFile)

		if spec != nil {
			for templateName := range loadedTemplates {
				template := loadedTemplates[templateName]
				codeGenerator := generator.New()
				template.UpdateOutputFilePath(codeOutputDirectory)
				codeGenerator.Set(spec, template)
				err = codeGenerator.Run()
				utils.CheckErrorForExit(err)

			}

		}

		fmt.Println("\nEverything looks fine.")
	},
}

func Execute() {
	if err := RootCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
