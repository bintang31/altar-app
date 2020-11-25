package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

func init() {
	workerCmd.SetUsageTemplate(workerUsage)
	rootCmd.AddCommand(workerCmd)
}

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Start worker service",
	Run:   workerHandler,
}

var workerHandler = func(cmd *cobra.Command, args []string) {

	if len(args) == 0 {
		cmd.Help()
		os.Exit(0)
	}
	option := args[0]

	log.Println(option)

}

var workerUsage = `
Run mobileloket worker

Usage:
altar worker [command]

Available Commands:
	start                Start altar worker
	stop                 Stop altar worker
// `
