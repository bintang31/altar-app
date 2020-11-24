package cmd

import (
	"altar-app/interfaces/routes"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start mobileloket http service",
	Run: func(cmd *cobra.Command, args []string) {
		routes.API()
	},
}
