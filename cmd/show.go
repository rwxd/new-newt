package cmd

import (
	"fmt"
	"log"

	"github.com/rwxd/new-newt/db"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show stuff",
}

var showDomainsCmd = &cobra.Command{
	Use:   "domains",
	Short: "show all domains that are being checked",
	Run: func(cmd *cobra.Command, args []string) {
		domains, err := db.Rdb.GetDomainsToCheck()
		if err != nil {
			log.Fatal(err)
		}

		for _, domain := range domains {
			fmt.Println(domain)
		}
	},
}

func init() {
	showCmd.AddCommand(showDomainsCmd)
}
