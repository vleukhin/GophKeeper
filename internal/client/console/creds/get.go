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
  -i, --id string Login id`,
	Run: func(cmd *cobra.Command, args []string) {
		client.GetApp().ShowCred(getLoginID)
	},
}

var getLoginID string //nolint:gochecknoglobals // cobra style guide

func init() {
	GetCred.Flags().StringVarP(&getLoginID, "id", "i", "", "Card id")

	if err := GetCred.MarkFlagRequired("id"); err != nil {
		log.Fatal(err)
	}
}
