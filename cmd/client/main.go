package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vleukhin/GophKeeper/internal/client"
	"github.com/vleukhin/GophKeeper/internal/client/api"
	"github.com/vleukhin/GophKeeper/internal/client/console/auth"
	"github.com/vleukhin/GophKeeper/internal/client/console/cards"
	"github.com/vleukhin/GophKeeper/internal/client/storage"
	config "github.com/vleukhin/GophKeeper/internal/config/client"
	"log"
)

var buildVersion = "N/A"
var buildDate = "N/A"
var buildCommit = "N/A"

var rootCmd = &cobra.Command{
	Short: "GophKeeper Client",
	Long:  `GothKeeper client stores your private data`,
	Run: func(cmd *cobra.Command, args []string) {
		printBuildInfo()
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initApp)
	commands := []*cobra.Command{
		auth.RegisterCmd,
		auth.LoginCmd,
		auth.LogoutCmd,
		cards.StoreCard,
		cards.GetCard,
		cards.DelCard,
	}

	rootCmd.AddCommand(commands...)
}

func initApp() {
	cfg := config.LoadConfig()

	app := client.GetApp()
	app.SetStorage(storage.NewMockStorage())
	app.SetConfig(cfg)
	app.SetAPIClient(api.NewHttpClient(cfg.Server.URL))
}

func printBuildInfo() {
	fmt.Println("GophKeeper Client")
	fmt.Println("----------------")
	fmt.Printf("Build version: %s\n", buildVersion)
	fmt.Printf("Build date: %s\n", buildDate)
	fmt.Printf("Build commit: %s\n", buildCommit)
	fmt.Println("----------------")
}
