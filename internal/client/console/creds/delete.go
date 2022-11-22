package creds

import (
	"log"

	"github.com/vleukhin/GophKeeper/internal/client"

	"github.com/spf13/cobra"
)

var DelCred = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "dellogin",
	Short: "Delete user login by id",
	Long: `
This command remove login
Usage: delcard -i \"login_id\" 
Flags:
  -i, --id string Card id
  -p, --password string   User password value.`,
	Run: func(cmd *cobra.Command, args []string) {
		client.GetApp().DelCred(userPassword, delLoginID)
	},
}

var delLoginID string //nolint:gochecknoglobals // cobra style guide

func init() {
	DelCred.Flags().StringVarP(&userPassword, "password", "p", "", "User password value.")
	DelCred.Flags().StringVarP(&delLoginID, "id", "i", "", "Card id")

	if err := DelCred.MarkFlagRequired("password"); err != nil {
		log.Fatal(err)
	}
	if err := DelCred.MarkFlagRequired("id"); err != nil {
		log.Fatal(err)
	}
}
