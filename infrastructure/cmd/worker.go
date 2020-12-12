package cmd

import (
	"altar-app/application/config"
	logger "altar-app/infrastructure/logger"
	"altar-app/infrastructure/queue/worker"
	"altar-app/interfaces/scheduler"
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
	conf := config.LoadAppConfig("amqp")
	queueName := conf.QueueName
	workerLimit := conf.WorkerLimit
	var consumer worker.AMQPConsumer = worker.AMQPConsumer{Consumer: make(map[string]worker.AMQPWorker)}
	switch option {
	case "start":
		interfaces.InitCronInfo()
		ok := consumer.StartWorker(queueName, workerLimit)
		if ok {
			logger.InfoLogHandler("Worker Running")
		}
	case "stop":
		consumer.StopWorker(queueName)
		logger.InfoLogHandler("Worker Stopped")
	}
}

var workerUsage = `
Run mobileloket worker

Usage:
altar worker [command]

Available Commands:
	start                Start altar worker
	stop                 Stop altar worker
// `
