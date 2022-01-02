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
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
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
