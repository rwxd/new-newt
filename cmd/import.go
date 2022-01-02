package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/rwxd/new-newt/db"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "import domains",
}

var importFileCmd = &cobra.Command{
	Use:   "file",
	Short: "import files with domains seperated by newline",
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			log.Printf("Importing domains from file: %s\n", arg)
			data, err := os.ReadFile(arg)
			if err != nil {
				log.Fatal("Cannot read file:", err)
			}

			s := string(data)
			lines := strings.Split(s, "\n")
			domains_to_add := make([]string, 0)
			for _, line := range lines {
				fmt.Printf("Adding \"%s\" to the list of domains to check\n", line)
				domains_to_add = append(domains_to_add, strings.ToLower(line))
			}
			err = db.Rdb.AddDomainsToCheck(domains_to_add)
			if err != nil {
				log.Println("error adding domains to list:", err)
			}
		}
	},
}

func init() {
	importCmd.AddCommand(importFileCmd)
}
