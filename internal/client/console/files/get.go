package files

import (
	"log"

	"github.com/vleukhin/GophKeeper/internal/client"

	"github.com/spf13/cobra"
)

var GetFile = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
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

var (
	fileID   string //nolint:gochecknoglobals // cobra style guide
	filePath string //nolint:gochecknoglobals // cobra style guide
)

func init() {
	GetFile.Flags().StringVarP(&fileID, "id", "i", "", "Text id")
	GetFile.Flags().StringVarP(&filePath, "file", "f", "", "File path")

	if err := GetFile.MarkFlagRequired("id"); err != nil {
		log.Fatal(err)
	}
	if err := GetFile.MarkFlagRequired("file"); err != nil {
		log.Fatal(err)
	}
}
