package config

import (
	"github.com/spf13/viper"
)

//Database : Struct Load Config
type Database struct {
	Driver            string
	Host              string
	User              string
	Password          string
	DBName            string
	QueueName         string
	WorkerLimit       int
	Port              string
	ReconnectRetry    int
	ReconnectInterval int64
	DebugMode         bool
}

//Loket : Struct Load Config Loket
type Loket struct {
	EndPoint  string
	User      string
	Password  string
	UserLoket string
}

// LoadAppConfig load database configuration
func LoadAppConfig(name string) Database {
	db := viper.Sub("database." + name)
	conf := Database{
		Driver:            db.GetString("driver"),
		Host:              db.GetString("host"),
		User:              db.GetString("user"),
		Password:          db.GetString("password"),
		DBName:            db.GetString("db_name"),
		QueueName:         db.GetString("queue_name"),
		WorkerLimit:       db.GetInt("worker_limit"),
		Port:              db.GetString("port"),
		ReconnectRetry:    db.GetInt("reconnect_retry"),
		ReconnectInterval: db.GetInt64("reconnect_interval"),
		DebugMode:         db.GetBool("debug"),
	}
	return conf
}

// LoadModuleLoketConfig load module loket configuration
func LoadModuleLoketConfig(name string) Loket {
	db := viper.Sub(name)
	conf := Loket{
		EndPoint:  db.GetString("endpoint"),
		User:      db.GetString("user"),
		Password:  db.GetString("password"),
		UserLoket: db.GetString("userloket"),
	}
	return conf
}
