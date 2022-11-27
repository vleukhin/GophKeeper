package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/vleukhin/GophKeeper/internal/client/console/creds"
	"github.com/vleukhin/GophKeeper/internal/client/console/notes"

	"github.com/spf13/cobra"

	"github.com/vleukhin/GophKeeper/internal/client"
	"github.com/vleukhin/GophKeeper/internal/client/api"
	"github.com/vleukhin/GophKeeper/internal/client/console/auth"
	"github.com/vleukhin/GophKeeper/internal/client/console/cards"
	"github.com/vleukhin/GophKeeper/internal/client/storage"
	"github.com/vleukhin/GophKeeper/internal/client/storage/postgres"
	config "github.com/vleukhin/GophKeeper/internal/config/client"
)

var buildVersion = "N/A" //nolint:gochecknoglobals
var buildDate = "N/A"    //nolint:gochecknoglobals
var buildCommit = "N/A"  //nolint:gochecknoglobals

func main() {
	var rootCmd = &cobra.Command{
		Short: "GophKeeper Client",
		Long:  `GothKeeper client stores your private data`,
		Run: func(cmd *cobra.Command, args []string) {
			printBuildInfo()
		},
	}

	cobra.OnInitialize(initApp)
	commands := []*cobra.Command{
		auth.NewRegisterCommand(),
		auth.NewLoginCommand(),
		auth.NewLogoutCommand(),
		creds.StoreCred,
		creds.GetCred,
		creds.DelCred,
		cards.StoreCard,
		cards.GetCard,
		cards.DelCard,
		notes.StoreNote,
		notes.GetNote,
		notes.DelNote,
	}

	rootCmd.AddCommand(commands...)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func initApp() {
	var repo storage.Repo
	var err error
	cfg := config.LoadConfig()

	switch cfg.Storage.Driver {
	case "mock":
		repo = storage.NewMockStorage()
	case "postgres":
		repo, err = postgres.NewPostgresStorage(cfg.Postgres.DSN, time.Second*5)
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("Unknown storage driver")
	}

	app := client.GetApp()
	app.SetStorage(repo)
	app.SetConfig(cfg)
	app.SetAPIClient(api.NewHTTPClient(cfg.Server.URL))
	err = app.InitDB(context.Background())
	if err != nil {
		log.Fatal(err.Error())
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
