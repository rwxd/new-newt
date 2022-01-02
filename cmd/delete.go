package cmd

import (
	"log"

	"github.com/rwxd/new-newt/db"
	"github.com/rwxd/new-newt/utils"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a domain from the check list",
	Run: func(cmd *cobra.Command, args []string) {
		domains, err := db.Rdb.GetDomainsToCheck()

		if err != nil {
			log.Fatal(err)
		}

		for _, arg := range args {
			if utils.StringInSlice(arg, domains) {
				log.Printf("Removing \"%s\" from the list of domains to check\n", arg)
				err = db.Rdb.DeleteDomainToCheck(arg)
				if err != nil {
					log.Println("Error adding domain to list:", err)
				}
			} else {
				log.Printf("\"%s\" is not in the list of domains to check\n", arg)
			}
		}

	},
}
