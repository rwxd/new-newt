package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "new-newt",
	Short: "Domain availability checker",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(crawlerCmd)
	rootCmd.AddCommand(webCmd)
	rootCmd.AddCommand(showCmd)
	rootCmd.AddCommand(clearCmd)
	rootCmd.AddCommand(importCmd)
}
