package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "city-stats",
		Short: "Fetches stats for a given city",
	}
)

func Execute() error {
	return rootCmd.Execute()
}
