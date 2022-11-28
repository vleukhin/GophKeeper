package files

import (
	"log"
	"os"

	"github.com/vleukhin/GophKeeper/internal/client"
	"github.com/vleukhin/GophKeeper/internal/client/console"
	"github.com/vleukhin/GophKeeper/internal/models"

	"github.com/spf13/cobra"
)

var StoreFile = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "storefile",
	Short: "Store file",
	Long: `
This command store file
Usage: storefile
Flags:
  -n, --name string       File human readable name
  -f, --file string       File path
  -h, --help              help for storefile
  -meta 				  Store meta data for entry
  example: -meta'[{"name":"some_meta","value":"some_meta_value"},{"name":"some_meta2","value":"some_meta_value2"}]'
  `,
	Run: func(cmd *cobra.Command, args []string) {
		client.GetApp().StoreFile(newFile)
	},
}

var (
	newFile models.File //nolint:gochecknoglobals // cobra style guide
)

func init() {
	var filePath string
	StoreFile.Flags().StringVarP(&newFile.Name, "title", "t", "", "Card title")
	StoreFile.Flags().StringVarP(&filePath, "file", "f", "", "File path")
	StoreFile.Flags().Var(&console.JSONFlag{Target: &newFile.Meta}, "meta", `Store meta fields for models`)

	if err := StoreFile.MarkFlagRequired("name"); err != nil {
		log.Fatal(err)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err.Error())
	}

	newFile.Content = content
}
