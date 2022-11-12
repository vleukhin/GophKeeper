package auth

import (
	"github.com/spf13/cobra"
	"github.com/vleukhin/GophKeeper/internal/client"
)

func NewRegisterCmd(app *client.App) *cobra.Command {
	return &cobra.Command{
		Use:   "register {login} {password}",
		Short: "Perform user registration",
		Long:  "Perform user registration",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			app.Register(args[0], args[1])
		},
	}
}
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
func NewLogoutCmd(app *client.App) *cobra.Command {
	return &cobra.Command{
		Use:   "logout",
		Short: "Logout user from system",
		Long:  "Logout user from system",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			app.Logout()
		},
	}
}
