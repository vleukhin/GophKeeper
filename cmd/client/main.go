package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vleukhin/GophKeeper/cmd/client/auth"
	"github.com/vleukhin/GophKeeper/internal/client"
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
	app := client.NewApp()
	// cobra.OnInitialize(initApp)
	commands := []*cobra.Command{
		auth.NewLoginCmd(app),
	}

	rootCmd.AddCommand(commands...)
}

func printBuildInfo() {
	fmt.Println("GophKeeper Client")
	fmt.Println("----------------")
	fmt.Printf("Build version: %s\n", buildVersion)
	fmt.Printf("Build date: %s\n", buildDate)
	fmt.Printf("Build commit: %s\n", buildCommit)
	fmt.Println("----------------")
}
