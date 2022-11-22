package cards

import (
	"log"

	"github.com/vleukhin/GophKeeper/internal/client"
	"github.com/vleukhin/GophKeeper/internal/client/console"
	"github.com/vleukhin/GophKeeper/internal/models"

	"github.com/spf13/cobra"
)

var StoreCard = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "storecard",
	Short: "Store newCard",
	Long: `
This command store newCard
Usage: storecard -p \"user_password\" 
Flags:
  -b, --bank string       Card issuer bank
  -c, --code string       Card code
  -h, --help              help for storecard
  -m, --month string      Card expiration month
  -n, --number string     Card number
  -o, --owner string      Card holder name
  -p, --password string   User password value.
  -t, --title string      Card title
  -y, --year string       Card expiration year
  -meta 				  Store meta data for entry
  example: -meta'[{"name":"some_meta","value":"some_meta_value"},{"name":"some_meta2","value":"some_meta_value2"}]'
  `,
	Run: func(cmd *cobra.Command, args []string) {
		client.GetApp().StoreCard(userPassword, &newCard)
	},
}

var (
	newCard      models.Card //nolint:gochecknoglobals // cobra style guide
	userPassword string      //nolint:gochecknoglobals // cobra style guide
)

func init() {
	StoreCard.Flags().StringVarP(&userPassword, "password", "p", "", "User password value.")
	StoreCard.Flags().StringVarP(&newCard.Name, "title", "t", "", "Card title")
	StoreCard.Flags().StringVarP(&newCard.Number, "number", "n", "", "Card number")
	StoreCard.Flags().StringVarP(&newCard.CardHolderName, "owner", "o", "", "Card holder name")
	StoreCard.Flags().StringVarP(&newCard.Bank, "bank", "b", "", "Card issuer bank")
	StoreCard.Flags().StringVarP(&newCard.SecurityCode, "code", "c", "", "Card code")
	StoreCard.Flags().StringVarP(&newCard.ExpirationMonth, "month", "m", "", "Card expiration month")
	StoreCard.Flags().StringVarP(&newCard.ExpirationYear, "year", "y", "", "Card expiration year")
	StoreCard.Flags().Var(&console.JSONFlag{Target: &newCard.Meta}, "meta", `Store meta fields for models`)

	if err := StoreCard.MarkFlagRequired("password"); err != nil {
		log.Fatal(err)
	}
	if err := StoreCard.MarkFlagRequired("title"); err != nil {
		log.Fatal(err)
	}
	if err := StoreCard.MarkFlagRequired("number"); err != nil {
		log.Fatal(err)
	}
	if err := StoreCard.MarkFlagRequired("password"); err != nil {
		log.Fatal(err)
	}
}
