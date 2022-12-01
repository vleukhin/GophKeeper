package cards

import (
	"log"

	"github.com/vleukhin/GophKeeper/internal/client"

	"github.com/spf13/cobra"
)

var GetCard = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "getcard",
	Short: "Show user newCard by id",
	Long: `
This command add newCard
Usage: getcard -i \"card_id\" 
Flags:
  -i, --id string Card id`,
	Run: func(cmd *cobra.Command, args []string) {
		client.GetApp().ShowCard(getCardID)
	},
}

var getCardID string //nolint:gochecknoglobals // cobra style guide

func init() {
	GetCard.Flags().StringVarP(&getCardID, "id", "i", "", "Card id")

	if err := GetCard.MarkFlagRequired("id"); err != nil {
		log.Fatal(err)
	}
}
