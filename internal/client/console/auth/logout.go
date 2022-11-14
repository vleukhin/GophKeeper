package auth

import (
	"github.com/spf13/cobra"
	"github.com/vleukhin/GophKeeper/internal/client"
)

var LogoutCmd = &cobra.Command{ //nolint:gochecknoglobals
	Use:   "logout",
	Short: "Performs user logout",
	Long:  "Performs user logout",
	Run: func(cmd *cobra.Command, args []string) {
		client.GetApp().Logout()
	},
}
