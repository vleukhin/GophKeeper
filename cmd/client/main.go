package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vleukhin/GophKeeper/internal/client/api"
	"github.com/vleukhin/GophKeeper/internal/client/console/auth"
	"github.com/vleukhin/GophKeeper/internal/client/console/cards"
	"github.com/vleukhin/GophKeeper/internal/client/core"
	"github.com/vleukhin/GophKeeper/internal/client/storage"
	"github.com/vleukhin/GophKeeper/internal/config/client"
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
		cards.DelCard,
	}

	rootCmd.AddCommand(commands...)
}

func initApp() {
	cfg := client.LoadConfig()

	app := core.GetApp()
	optFuncs := []core.OptFunc{
		core.SetAPIClient(api.NewHttpClient(cfg.Server.URL)),
		core.SetConfig(cfg),
		core.SetRepo(storage.NewMockStorage()),
	}

	for _, f := range optFuncs {
		f(app)
	}
}

func printBuildInfo() {
	fmt.Println("GophKeeper Client")
	fmt.Println("----------------")
	fmt.Printf("Build version: %s\n", buildVersion)
	fmt.Printf("Build date: %s\n", buildDate)
	fmt.Printf("Build commit: %s\n", buildCommit)
	fmt.Println("----------------")
}
