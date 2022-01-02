package cmd

import (
	"fmt"
	"log"

	"github.com/rwxd/new-newt/db"
	"github.com/rwxd/new-newt/utils"
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

		for _, domain := range domains {
			fmt.Println(domain)
		}
	},
}

func init() {
	showCmd.AddCommand(showDomainsCmd)
}
