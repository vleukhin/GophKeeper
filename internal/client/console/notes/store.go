package notes

import (
	"log"

	"github.com/vleukhin/GophKeeper/internal/client"
	"github.com/vleukhin/GophKeeper/internal/client/console"
	"github.com/vleukhin/GophKeeper/internal/models"

	"github.com/spf13/cobra"
)

var StoreNote = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "addnote",
	Short: "Add note",
	Long: `
This command add user note
Usage: addnote
Flags:
  -h, --help            help for addlogin
  -n, --note string     User note  
  -meta 				Add meta data for entiry
  example: -meta'[{"name":"some_meta","value":"some_meta_value"},{"name":"some_meta2","value":"some_meta_value2"}]'
  `,
	Run: func(cmd *cobra.Command, args []string) {
		client.GetApp().StoreNote(&newNote)
	},
}

var (
	newNote      models.Note //nolint:gochecknoglobals // cobra style guide
	userPassword string      //nolint:gochecknoglobals // cobra style guide
)

func init() {
	StoreNote.Flags().StringVarP(&newNote.Name, "title", "t", "", "Login title")
	StoreNote.Flags().StringVarP(&newNote.Text, "note", "n", "", "User note")
	StoreNote.Flags().Var(&console.JSONFlag{Target: &newNote.Meta}, "meta", `Add meta fields for models`)

	if err := StoreNote.MarkFlagRequired("title"); err != nil {
		log.Fatal(err)
	}
}
