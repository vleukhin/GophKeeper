package app

import (
	"github.com/vleukhin/GophKeeper/internal/client"
	"log"

	"github.com/spf13/cobra"
	"github.com/vleukhin/GophKeeper/internal/entity"
	utils "github.com/vleukhin/GophKeeper/internal/utils/client"
)

var StoreCard = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "storecard",
	Short: "Store card",
	Long: `
This command store card
Usage: storecard -p \"user_password\" 
Flags:
  -b, --brand string      Card brand
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
		client.GetApp().AddCard(userPassword, &cardForStoreiting)
	},
}

var (
	cardForStoreiting entity.Card //nolint:gochecknoglobals // cobra style guide
	userPassword      string      //nolint:gochecknoglobals // cobra style guide
)

func init() {
	StoreCard.Flags().StringVarP(&userPassword, "password", "p", "", "User password value.")
	StoreCard.Flags().StringVarP(&cardForStoreiting.Name, "title", "t", "", "Card title")
	StoreCard.Flags().StringVarP(&cardForStoreiting.Number, "number", "n", "", "Card namber")
	StoreCard.Flags().StringVarP(&cardForStoreiting.CardHolderName, "owner", "o", "", "Card holder name")
	StoreCard.Flags().StringVarP(&cardForStoreiting.Brand, "brand", "b", "", "Card brand")
	StoreCard.Flags().StringVarP(&cardForStoreiting.SecurityCode, "code", "c", "", "Card code")
	StoreCard.Flags().StringVarP(&cardForStoreiting.ExpirationMonth, "month", "m", "", "Card expiration month")
	StoreCard.Flags().StringVarP(&cardForStoreiting.ExpirationYear, "year", "y", "", "Card expiration year")
	StoreCard.Flags().Var(&utils.JSONFlag{Target: &cardForStoreiting.Meta}, "meta", `Store meta fields for models`)

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
