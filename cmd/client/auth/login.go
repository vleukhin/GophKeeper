package auth

import (
	"github.com/spf13/cobra"
	"github.com/vleukhin/GophKeeper/internal/client/core"
	"github.com/vleukhin/GophKeeper/internal/models"
)

var LoginCmd = &cobra.Command{ //nolint:gochecknoglobals
	Use:   "login {login} {password}",
	Short: "Performs user login",
	Long:  "Performs user login",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		core.GetApp().Login(&models.User{
			Email:    args[0],
			Password: args[1],
		})
	},
}
