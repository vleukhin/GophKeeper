package cards

import (
	"log"

	"github.com/vleukhin/GophKeeper/internal/client"

	"github.com/spf13/cobra"
)

var DelCard = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "delcard",
	Short: "Delete user newCard by id",
	Long: `
This command remove newCard
Usage: delcard -i \"card_id\" 
Flags:
  -i, --id string Card id`,
	Run: func(cmd *cobra.Command, args []string) {
		client.GetApp().DelCard(delCardID)
	},
}

var delCardID string //nolint:gochecknoglobals // cobra style guide

func init() {
	DelCard.Flags().StringVarP(&delCardID, "id", "i", "", "Card id")

	if err := DelCard.MarkFlagRequired("id"); err != nil {
		log.Fatal(err)
	}
}
