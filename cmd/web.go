package cmd

import (
	"log"
	"strconv"

	"github.com/rwxd/new-newt/utils"
	"github.com/rwxd/new-newt/web"
	"github.com/spf13/cobra"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "serves the web interface",
	Run: func(cmd *cobra.Command, args []string) {
		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			log.Fatal(err)
		}
		config, err := utils.LoadConfig(".")
		if err != nil {
			log.Fatal("Cannot load config:", err)
		}

		web.WebServer(strconv.Itoa(port), config)
	},
}

func init() {
	webCmd.Flags().IntP("port", "p", 8080, "port to serve the web interface on")
}
