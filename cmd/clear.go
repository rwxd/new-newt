package cmd

import (
	"log"

	"github.com/rwxd/new-newt/db"
	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "clear list of domains to check",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Clearing list of domains to check")
		err := db.Rdb.ClearDomains()
		if err != nil {
			log.Fatal(err)
		}
	},
}
