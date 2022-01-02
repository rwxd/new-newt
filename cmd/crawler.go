package cmd

import (
	"log"
	"sync"
	"time"

	"github.com/rwxd/new-newt/db"
	"github.com/rwxd/new-newt/whois"
	"github.com/spf13/cobra"
)

var crawlerCmd = &cobra.Command{
	Use:   "crawler",
	Short: "runs the crawler in an infinity loop",
	Run: func(cmd *cobra.Command, args []string) {
		interval, _ := cmd.Flags().GetInt("interval")
		checkInterval := time.Duration(interval) * time.Minute
		concurrent, _ := cmd.Flags().GetInt("concurrent")

		log.Println("Starting the application...")
		log.Println("Check interval:", checkInterval)
		log.Println("Concurrent checks:", concurrent)

		whoisClient := whois.NewWhoisClient()
		var wg sync.WaitGroup
		guard := make(chan struct{}, concurrent)

		for {
			domains, err := db.Rdb.GetDomainsToCheck()
			if err != nil {
				log.Fatal(err)
			}
			if len(domains) == 0 {
				log.Println("No domains to check")
			}
			for _, domain := range domains {
				guard <- struct{}{} // would block if guard channel is already filled
				wg.Add(1)
				go func(domain string) {
					defer wg.Done()

					status, err := db.Rdb.GetDomainStatus(domain)
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
						db.Rdb.SetDomainStatus(domain, status)
					} else {
						log.Printf("Domain %s is not available\n", domain)
						status := db.NewDomainStatus(false, time)
						db.Rdb.SetDomainStatus(domain, status)
					}
					<-guard // release guard
				}(domain)
			}
			time.Sleep(60 * time.Second)
		}
	},
}

func init() {
	crawlerCmd.Flags().IntP("interval", "i", 60, "interval between checks in minutes")
	crawlerCmd.Flags().IntP("concurrent", "c", 10, "number of concurrent checks")
}
