package auth

import (
	"github.com/spf13/cobra"
	"github.com/vleukhin/GophKeeper/internal/client"
)

func NewLoginCmd(app *client.App) *cobra.Command {
	return &cobra.Command{
		Use:   "login {login} {password}",
		Short: "Perform user login",
		Long:  "Perform user login",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			app.Login(args[0], args[1])
		},
	}
}
