package config

import (
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

// Config ...
type Config struct {
	Port                 string `envconfig:"port"`
	MysqlHost            string `envconfig:"mysql_host" default:"localhost"`
	MysqlUsername        string `envconfig:"mysql_username" default:"root"`
	MysqlPassword        string `envconfig:"mysql_password" default:"root-is-not-used"`
	MysqlConnectionLimit int    `envconfig:"mysql_connection_limit" default:"40"`
	MysqlDatabase        string `envconfig:"mysql_database" default:"ecommerce"`
}

var conf Config
var once sync.Once

// Get returns the singleton config instance
func Get() Config {
	once.Do(func() {
		err := envconfig.Process("", &conf)
		if err != nil {
			log.Fatal("Can't load config: ", err)
		}
	})

	return conf
}
