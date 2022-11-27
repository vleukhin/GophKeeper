package auth

import (
	"github.com/spf13/cobra"

	"github.com/vleukhin/GophKeeper/internal/client"
	"github.com/vleukhin/GophKeeper/internal/models"
)

func NewRegisterCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "register {login} {password}",
		Short: "Performs user registration",
		Long:  "Performs user registration",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			client.GetApp().Register(models.User{
				Name:     args[0],
				Password: args[1],
			})
		},
	}
}
