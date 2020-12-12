package producer

import (
	"altar-app/application/config"
	log "altar-app/infrastructure/logger"
	"github.com/gofort/dispatcher"
	"github.com/spf13/viper"
)

type AMQPProducer struct {
	Server map[string]*dispatcher.Server
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Error().Msgf("%s: %s", msg, err)
	}
}

//CreateQueue : function to create queue
func (prod *AMQPProducer) CreateQueue(queueName string) {
	conf := config.LoadAppConfig("amqp")
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
		log.Fatal().Msgf("Error Start Worker: %s", err.Error())
		return
	}
	prod.Server[queueName] = server
}

func (prod *AMQPProducer) getServer(queueName string) *dispatcher.Server {
	server := prod.Server[queueName]
	if server == nil {
		prod.CreateQueue(queueName)
		server = prod.Server[queueName]
	}
	return server
}

func (prod *AMQPProducer) CreateItem(queueName string, payload string) {
	if viper.Get("env") == "testing" {
		return
	}
	server := prod.getServer(queueName)
	task := &dispatcher.Task{
		Name: queueName,
		Args: []dispatcher.TaskArgument{
			{
				Type:  "string",
				Value: payload,
			},
		},
	}
	// Here we sending task to a queue
	if err := server.Publish(task); err != nil {
		log.Fatal().Msgf("Error Create Queue: %s", err.Error())
		return
	}
}

var Producer AMQPProducer = AMQPProducer{make(map[string]*dispatcher.Server)}
