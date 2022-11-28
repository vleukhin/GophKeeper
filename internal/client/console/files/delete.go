package files

import (
	"log"

	"github.com/vleukhin/GophKeeper/internal/client"

	"github.com/spf13/cobra"
)

var DelFile = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
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

var delFileID string //nolint:gochecknoglobals // cobra style guide

func init() {
	DelFile.Flags().StringVarP(&delFileID, "id", "i", "", "File id")

	if err := DelFile.MarkFlagRequired("id"); err != nil {
		log.Fatal(err)
	}
}
