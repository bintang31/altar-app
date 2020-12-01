package worker

import (
	"altar-app/application/config"
	log "altar-app/infrastructure/logger"
	"altar-app/infrastructure/queue/task"
	"github.com/gofort/dispatcher"
	logs "log"
	"os"
	"os/signal"
	"time"
)

type AMQPWorker struct {
	WorkerName string
	Server     *dispatcher.Server
	Worker     *dispatcher.Worker
	Limit      int
	IsActive   bool
	Status     string
	TaskName   string
	//Function   interface{}
}

type AMQPConsumer struct {
	Consumer map[string]AMQPWorker
}

func (cons *AMQPConsumer) CreateServer(queueName string) {
	conf := config.LoadAppConfig("aqmp")
	amqp_url := "amqp://" + conf.User + ":" + conf.Password + "@" + conf.Host + ":" + conf.Port
	cfg := dispatcher.ServerConfig{
		AMQPConnectionString:        amqp_url,
		ReconnectionRetries:         conf.ReconnectRetry,
		ReconnectionIntervalSeconds: conf.ReconnectInterval,
		DebugMode:                   conf.DebugMode, // enables extended logging
		Exchange:                    queueName,
		InitQueues: []dispatcher.Queue{ // creates queues and binding keys if they are not created already
			{
				Name:        queueName,
				BindingKeys: []string{queueName},
			},
		},
		DefaultRoutingKey: queueName, // default routing key which is used for publishing messages
	}

	// This function creates new server (server consists of AMQP connection and publisher which sends tasks)
	server, _, err := dispatcher.NewServer(&cfg)
	if err != nil {
		log.Fatal().Msgf("Error Create Worker: %s", err.Error())
		logs.Println("Error Create Worker: ", err.Error())
		return
	}
	cons.Consumer[queueName] = AMQPWorker{queueName, server, nil, 0, false, "SHUTDOWN", "default"}
}

func (cons *AMQPConsumer) getServer(queueName string) *dispatcher.Server {
	server := cons.Consumer[queueName].Server
	if server == nil {
		cons.CreateServer(queueName)
		server = cons.Consumer[queueName].Server
	}
	return server
}

func (cons *AMQPConsumer) StartWorker(queueName string, limit int) bool {
	server := cons.getServer(queueName)
	if limit <= 0 {
		limit = 5
	}

	consumer := cons.Consumer[queueName]
	if limit == consumer.Limit {
		if worker, err := server.GetWorkerByName(consumer.WorkerName); err == nil {
			if err := worker.Start(server); err != nil {
				return false
			}
			consumer.Limit = limit
			consumer.Worker = worker
			consumer.IsActive = true
			consumer.Status = "RUNNING"
			cons.Consumer[queueName] = consumer
			return true
		}
	}

	workerName := "worker_" + queueName + "_" + time.Now().Format("0102150405")
	// Basic worker configuration
	workercfg := dispatcher.WorkerConfig{
		Queue: queueName,
		Name:  workerName,
		Limit: limit,
	}

	tasks := make(map[string]dispatcher.TaskConfig)

	// Task configuration where we pass function which will be executed by this worker when this task will be received
	tasks[queueName] = dispatcher.TaskConfig{
		Function: task.RunTask,
	}

	// This function creates worker, but he won't start to consume messages here
	worker, err := server.NewWorker(&workercfg, tasks)
	if err != nil {
		log.Fatal().Msgf("Error Start Worker: %s", err.Error())
		logs.Println("Error Start Worker: ", err.Error())
		return false
	}
	consumer = cons.Consumer[queueName]
	consumer.WorkerName = workerName
	consumer.Limit = limit
	consumer.Worker = worker
	consumer.IsActive = true
	consumer.Status = "RUNNING"
	cons.Consumer[queueName] = consumer

	// Here we start worker consuming
	if err := cons.Consumer[queueName].Worker.Start(server); err != nil {
		log.Fatal().Msgf("Error Start Worker: %s", err.Error())
		logs.Println("Error Start Worker: ", err.Error())
		consumer.IsActive = false
		consumer.Status = "SHUTDOWN"
		cons.Consumer[queueName] = consumer
		return false
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	server.Close()

	return true
}

func (cons *AMQPConsumer) StopWorker(queueName string) {
	// Here we stop worker consuming
	worker := cons.Consumer[queueName].Worker
	if worker != nil {
		logs.Println("worker found")
		// Update worker state
		consumer := cons.Consumer[queueName]
		consumer.IsActive = false
		consumer.Status = "STOPPING"
		cons.Consumer[queueName] = consumer
		worker.Close()
		// Update worker state
		consumer.Limit = 0
		consumer.Worker = worker
		consumer.IsActive = false
		consumer.Status = "SHUTDOWN"
		//cons.Consumer[queueName] = consumer
	}
}
