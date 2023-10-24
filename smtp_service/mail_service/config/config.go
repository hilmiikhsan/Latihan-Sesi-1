package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App  AppConfig  `yaml:"app"`
	Mail MailConfig `yaml:"mail"`
}

type AppConfig struct {
	Port string `yaml:"port"`
}

type MailConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Email    string `yaml:"email"`
	Password string `yaml:"password"`
}

func LoadConfigYaml(filename string) (conf Config, err error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return conf, err
	}

	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}
