package cmd

import (
	"log"

	"github.com/rwxd/new-newt/utils"
	"github.com/rwxd/new-newt/web"
	"github.com/spf13/cobra"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "serves the web interface",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := utils.LoadConfig(".")
		if err != nil {
			log.Fatal("Cannot load config:", err)
		}

		web.WebServer("8080", config)
	},
}
