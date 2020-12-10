package config

import (
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

// Config ...
type Config struct {
	Port             string `envconfig:"port"`
	PostgresHost     string `envconfig:"mysql_host" default:"localhost"`
	PostgresUsername string `envconfig:"mysql_username" default:"ecommerce_user"`
	PostgresPassword string `envconfig:"mysql_password" default:"ecommerce_password"`
	PostgresDatabase string `envconfig:"mysql_database" default:"ecommerce"`
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
