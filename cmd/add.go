package cmd

import (
	"log"
	"strings"

	"github.com/rwxd/new-newt/db"
	"github.com/rwxd/new-newt/utils"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a domain to the check list",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := utils.LoadConfig(".")
		if err != nil {
			log.Fatal("cannot load config:", err)
		}

		redisClient := db.NewRedisClient(config.RedisHost)

		if err != nil {
			log.Fatal("redis connection failed:", err)
		}

		domains, err := redisClient.GetDomainsToCheck()

		if err != nil {
			log.Fatal(err)
		}

		for _, arg := range args {
			if !utils.StringInSlice(arg, domains) {
				if !strings.Contains(arg, ".") {
					log.Printf("\"%s\" is not a valid domain name", arg)
					continue
				}

				log.Printf("Adding \"%s\" to the list of domains to check\n", arg)
				err = redisClient.AddDomainToCheck(arg)
				if err != nil {
					log.Println("error adding domain to list:", err)
				}
			}
		}

	},
}

func init() {
}
