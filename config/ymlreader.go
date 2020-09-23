package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type instanceConfig struct {
	Name     string `yaml:"Name"`
	Address  string `yaml:"Address"`
	Phone    int    `yaml:"Phone"`
	WFH      bool
	Database string `yaml:"Database"`
}

func (c *instanceConfig) Parse(data []byte) error {
	return yaml.Unmarshal(data, c)
}

func GetConfig() instanceConfig {
	return config
}

var config instanceConfig

func Init() {
	data, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		log.Fatal(err)
	}

	if err := config.Parse(data); err != nil {
		log.Fatal(err)
	}
}
