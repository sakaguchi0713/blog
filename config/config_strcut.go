package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	DB struct {
		Adapter  string `yaml:"adapter"`
		Encoding string `yaml:"encoding"`
		Database string `yaml:"database"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	}
}

func NewConfig() (*Config, error) {
	buf, err := ioutil.ReadFile("..//config.yml")
	if err != nil {
		return nil, err
	}
	var c Config
	err = yaml.Unmarshal(buf, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
