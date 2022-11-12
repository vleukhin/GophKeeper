package client

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var buildVersion = "N/A"
var buildDate = "N/A"
var buildCommit = "N/A"

func Run() {
	cmd := cobra.Command{
		Use:   "GophKeeper Client",
		Short: "GophKeeper Client",
		Long:  `GothKeeper client stores your private data`,
		Run: func(cmd *cobra.Command, args []string) {
			printBuildInfo()
		},
	}

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
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
