package creds

import (
	"log"

	"github.com/vleukhin/GophKeeper/internal/client"
	"github.com/vleukhin/GophKeeper/internal/client/console"
	"github.com/vleukhin/GophKeeper/internal/models"

	"github.com/spf13/cobra"
)

var StoreCred = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "addcred",
	Short: "Add newCred",
	Long: `
This command add logit for site
Usage: addcred -p \"user_password\" 
Flags:
  -h, --help              help for addcred
  -l, --login string      Site login
  -p, --password string   User password value.
  -s, --secret string     Site password|secret
  -t, --title string      Cred title
  -u, --uri string        Site endpoint  
  -meta 				  Add meta data for entry
  example: -meta'[{"name":"some_meta","value":"some_meta_value"},{"name":"some_meta2","value":"some_meta_value2"}]'
  `,
	Run: func(cmd *cobra.Command, args []string) {
		client.GetApp().StoreCred(userPassword, &newCred)
	},
}

var (
	newCred      models.Cred //nolint:gochecknoglobals // cobra style guide
	userPassword string      //nolint:gochecknoglobals // cobra style guide
)

func init() {
	StoreCred.Flags().StringVarP(&userPassword, "password", "p", "", "User password value.")

	StoreCred.Flags().StringVarP(&newCred.Name, "title", "t", "", "Cred title")
	StoreCred.Flags().StringVarP(&newCred.Login, "newCred", "l", "", "Site newCred")
	StoreCred.Flags().StringVarP(&newCred.Password, "secret", "s", "", "Site password|secret")
	StoreCred.Flags().StringVarP(&newCred.URI, "uri", "u", "", "Site endpoint")
	StoreCred.Flags().Var(&console.JSONFlag{Target: &newCred.Meta}, "meta", `Add meta fields for models`)

	if err := StoreCred.MarkFlagRequired("password"); err != nil {
		log.Fatal(err)
	}
	if err := StoreCred.MarkFlagRequired("title"); err != nil {
		log.Fatal(err)
	}
}
