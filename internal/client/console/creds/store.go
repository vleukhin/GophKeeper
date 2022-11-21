package app

import (
	"github.com/vleukhin/GophKeeper/internal/client"
	"github.com/vleukhin/GophKeeper/internal/client/console"
	"github.com/vleukhin/GophKeeper/internal/models"
	"log"

	"github.com/spf13/cobra"
)

var AddCred = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "addcred",
	Short: "Add cred",
	Long: `
This command add logit for site
Usage: addcred -p \"user_password\" 
Flags:
  -h, --help              help for addcred
  -l, --login string      Site login
  -p, --password string   User password value.
  -s, --secret string     Site password|secret
  -t, --title string      Cred title
  -u, --uri string        Site endloint  
  -meta 				  Add meta data for entiry
  example: -meta'[{"name":"some_meta","value":"some_meta_value"},{"name":"some_meta2","value":"some_meta_value2"}]'
  `,
	Run: func(cmd *cobra.Command, args []string) {
		client.GetApp().StoreCred(userPassword, &cred)
	},
}

var (
	cred         models.Cred //nolint:gochecknoglobals // cobra style guide
	userPassword string      //nolint:gochecknoglobals // cobra style guide
)

func init() {
	AddCred.Flags().StringVarP(&userPassword, "password", "p", "", "User password value.")

	AddCred.Flags().StringVarP(&cred.Name, "title", "t", "", "Cred title")
	AddCred.Flags().StringVarP(&cred.Login, "cred", "l", "", "Site cred")
	AddCred.Flags().StringVarP(&cred.Password, "secret", "s", "", "Site password|secret")
	AddCred.Flags().StringVarP(&cred.URI, "uri", "u", "", "Site endpoint")
	AddCred.Flags().Var(&console.JSONFlag{Target: &cred.Meta}, "meta", `Add meta fields for models`)

	if err := AddCred.MarkFlagRequired("password"); err != nil {
		log.Fatal(err)
	}
	if err := AddCred.MarkFlagRequired("title"); err != nil {
		log.Fatal(err)
	}
}
