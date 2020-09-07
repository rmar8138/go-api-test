package config

import (
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

// Configuration global config variable
var Configuration *Specification

func init() {
	c, err := LoadEnv()
	if err != nil {
		log.Fatal(err)
	}
	Configuration = c
}

// Specification for basic config
type Specification struct {
	Name     string `envconfig:"SERVICE_NAME" default:"go-api-test"`
	Debug    bool   `envconfig:"DEBUG" default:"true"`
	LogLevel string `envconfig:"LOG_LEVEL" default:"debug"`
	Port     string `envconfig:"PORT" default:"8000"`
}

// LoadEnv load env variables
func LoadEnv() (*Specification, error) {
	c := &Specification{}
	err := envconfig.Process("", c)
	return c, err
}
