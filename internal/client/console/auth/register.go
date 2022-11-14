package auth

import (
	"github.com/spf13/cobra"
	"github.com/vleukhin/GophKeeper/internal/client/core"
	"github.com/vleukhin/GophKeeper/internal/models"
)

var RegisterCmd = &cobra.Command{ //nolint:gochecknoglobals
	Use:   "register {login} {password}",
	Short: "Performs user registration",
	Long:  "Performs user registration",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		core.GetApp().Register(&models.User{
			Email:    args[0],
			Password: args[1],
		})
	},
}