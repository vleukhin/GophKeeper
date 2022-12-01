package auth

import (
	"github.com/spf13/cobra"

	"github.com/vleukhin/GophKeeper/internal/client"
)

func NewLoginCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "login {login} {password}",
		Short: "Performs user login",
		Long:  "Performs user login",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			client.GetApp().Login(args[0], args[1])
		},
	}
}
