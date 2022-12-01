package files

import (
	"log"
	"os"
	"path/filepath"

	"github.com/vleukhin/GophKeeper/internal/client/console"

	"github.com/vleukhin/GophKeeper/internal/client"
	"github.com/vleukhin/GophKeeper/internal/models"

	"github.com/spf13/cobra"
)

func NewStoreCommand() *cobra.Command {
	var filePath string
	var newFile models.File

	cmd := &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
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
			content, err := os.ReadFile(filePath)
			if err != nil {
				log.Fatal(err.Error())
			}

			newFile.Content = content
			newFile.FileName = filepath.Base(filePath)
			client.GetApp().StoreFile(newFile)
		},
	}

	cmd.Flags().StringVarP(&newFile.Name, "name", "n", "", "File human readable name")
	cmd.Flags().StringVarP(&filePath, "file", "f", "", "File path")
	cmd.Flags().Var(&console.JSONFlag{Target: &newFile.Meta}, "meta", `Store meta fields for models`)
	if err := cmd.MarkFlagRequired("name"); err != nil {
		log.Fatal(err)
	}

	return cmd
}

func NewGetCommand() *cobra.Command {
	var fileID string
	var filePath string
	cmd := &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
		Use:   "getfile",
		Short: "Show user file by id",
		Long: `
This command gets file
Usage: getfile -i \"file_id\" 
Flags:
  -i, --id string   File id
  -f, --file string File path`,
		Run: func(cmd *cobra.Command, args []string) {
			client.GetApp().GetFile(fileID, filePath)
		},
	}

	cmd.Flags().StringVarP(&fileID, "id", "i", "", "Text id")
	cmd.Flags().StringVarP(&filePath, "file", "f", "", "File path")

	if err := cmd.MarkFlagRequired("id"); err != nil {
		log.Fatal(err)
	}
	if err := cmd.MarkFlagRequired("file"); err != nil {
		log.Fatal(err)
	}

	return cmd
}

func NewDeleteCommand() *cobra.Command {
	var delFileID string

	cmd := &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
		Use:   "delfile",
		Short: "Delete file by id",
		Long: `
This command remove file	
Usage: delfile -i \"file_id\" 
Flags:
  -i, --id string File id`,
		Run: func(cmd *cobra.Command, args []string) {
			client.GetApp().DelFile(delFileID)
		},
	}

	cmd.Flags().StringVarP(&delFileID, "id", "i", "", "File id")
	if err := cmd.MarkFlagRequired("id"); err != nil {
		log.Fatal(err)
	}

	return cmd

}
