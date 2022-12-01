package notes

import (
	"log"

	"github.com/vleukhin/GophKeeper/internal/client"

	"github.com/spf13/cobra"
)

var DelNote = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "delnote",
	Short: "Delete note by id",
	Long: `
This command remove note	
Usage: delnote -i \"note_id\" 
Flags:
  -i, --id string Card id`,
	Run: func(cmd *cobra.Command, args []string) {
		client.GetApp().DelNote(delNoteID)
	},
}

var delNoteID string //nolint:gochecknoglobals // cobra style guide

func init() {
	DelNote.Flags().StringVarP(&delNoteID, "id", "i", "", "Text id")

	if err := DelNote.MarkFlagRequired("id"); err != nil {
		log.Fatal(err)
	}
}
