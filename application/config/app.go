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
	Port              string
	ReconnectRetry    int
	ReconnectInterval int64
	DebugMode         bool
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
		Port:              db.GetString("port"),
		ReconnectRetry:    db.GetInt("reconnect_retry"),
		ReconnectInterval: db.GetInt64("reconnect_interval"),
		DebugMode:         db.GetBool("debug"),
	}
	return conf
}
