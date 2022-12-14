package notes

import (
	"log"

	"github.com/vleukhin/GophKeeper/internal/client"

	"github.com/spf13/cobra"
)

var GetNote = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "getnote",
	Short: "Show user note by id",
	Long: `
This command show user note
Usage: getnote -i \"note_id\" 
Flags:
  -i, --id string Text id`,
	Run: func(cmd *cobra.Command, args []string) {
		client.GetApp().ShowNote(getNoteID)
	},
}

var getNoteID string //nolint:gochecknoglobals // cobra style guide

func init() {
	GetNote.Flags().StringVarP(&getNoteID, "id", "i", "", "Text id")

	if err := GetNote.MarkFlagRequired("id"); err != nil {
		log.Fatal(err)
	}
}
