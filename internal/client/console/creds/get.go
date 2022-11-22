package creds

import (
	"log"

	"github.com/vleukhin/GophKeeper/internal/client"

	"github.com/spf13/cobra"
)

var GetCred = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "getlogin",
	Short: "Show user login by id",
	Long: `
This command getlogin
Usage: getlogin -i \"login_id\" 
Flags:
  -i, --id string Login id
  -p, --password string   User password value.`,
	Run: func(cmd *cobra.Command, args []string) {
		client.GetApp().ShowCred(userPassword, getLoginID)
	},
}

var getLoginID string //nolint:gochecknoglobals // cobra style guide

func init() {
	GetCred.Flags().StringVarP(&userPassword, "password", "p", "", "User password value.")
	GetCred.Flags().StringVarP(&getLoginID, "id", "i", "", "Card id")

	if err := GetCred.MarkFlagRequired("password"); err != nil {
		log.Fatal(err)
	}
	if err := GetCred.MarkFlagRequired("id"); err != nil {
		log.Fatal(err)
	}
}
