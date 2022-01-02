package cmd

import (
	"log"

	"github.com/rwxd/new-newt/db"
	"github.com/rwxd/new-newt/utils"
	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "clear list of domains to check",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := utils.LoadConfig(".")
		if err != nil {
			log.Fatal("cannot load config:", err)
		}

		redisClient := db.NewRedisClient(config.RedisHost)

		if err != nil {
			log.Fatal("redis connection failed:", err)
		}

		log.Println("Clearing list of domains to check")
		err = redisClient.ClearDomains()
		if err != nil {
			log.Fatal(err)
		}
	},
}
