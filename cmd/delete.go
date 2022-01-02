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
		config, err := utils.LoadConfig(".")
		if err != nil {
			log.Fatal("Cannot load config:", err)
		}

		redisClient := db.NewRedisClient(config.RedisHost)

		if err != nil {
			log.Fatal("Redis connection failed:", err)
		}

		domains, err := redisClient.GetDomainsToCheck()

		if err != nil {
			log.Fatal(err)
		}

		for _, arg := range args {
			if utils.StringInSlice(arg, domains) {
				log.Printf("Removing \"%s\" from the list of domains to check\n", arg)
				err = redisClient.DeleteDomainToCheck(arg)
				if err != nil {
					log.Println("Error adding domain to list:", err)
				}
			} else {
				log.Printf("\"%s\" is not in the list of domains to check\n", arg)
			}
		}

	},
}
