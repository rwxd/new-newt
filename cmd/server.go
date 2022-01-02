package cmd

import (
	"log"
	"sync"
	"time"

	"github.com/rwxd/new-newt/db"
	"github.com/rwxd/new-newt/utils"
	"github.com/rwxd/new-newt/whois"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "serve the server",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Starting the application...")
		config, err := utils.LoadConfig(".")
		if err != nil {
			log.Fatal("cannot load config:", err)
		}

		whoisClient := whois.NewWhoisClient()
		redisClient := db.NewRedisClient(config.RedisHost)

		if err != nil {
			log.Fatal("redis connection failed:", err)
		}

		var wg sync.WaitGroup
		maxGoroutines := 1
		guard := make(chan struct{}, maxGoroutines)

		checkInterval := time.Minute * 60 * 24

		for {
			domains, err := redisClient.GetDomainsToCheck()
			if err != nil {
				log.Fatal(err)
			}
			for _, domain := range domains {
				guard <- struct{}{} // would block if guard channel is already filled
				wg.Add(1)
				go func(domain string) {
					defer wg.Done()

					status, err := redisClient.GetDomainStatus(domain)
					if err != nil {
						log.Printf("Domain \"%s\" not found in database\n", domain)
					} else {
						if status.LastSearched.Before(time.Now().Add(-checkInterval)) {
							log.Printf("Domain \"%s\" was last checked %s ago, checking again\n", domain, time.Since(status.LastSearched))
						} else {
							log.Printf("Domain \"%s\" was last checked %s ago, skipping\n", domain, time.Since(status.LastSearched))
							<-guard // release guard channel
							return
						}
					}

					log.Printf("Querying whois for \"%s\"\n", domain)
					response, err := whoisClient.Lookup(domain)
					if err != nil {
						log.Printf("Error querying whois for %s: %s\n", domain, err)
					}
					time := time.Now()
					if response.Available {
						log.Printf("Domain %s is available\n", domain)
						status := db.NewDomainStatus(true, time)
						redisClient.SetDomainStatus(domain, status)
					} else {
						log.Printf("Domain %s is not available\n", domain)
						status := db.NewDomainStatus(false, time)
						redisClient.SetDomainStatus(domain, status)
					}
					<-guard // release guard
				}(domain)
			}
			time.Sleep(60 * time.Second)
		}
	},
}
