package main

import (
	"github.com/go-yaml/yaml"
	"os"
)

type Config struct {
	Endpoints map[string]Endpoint `yaml:"endpoints"`
}

func LoadConfig(file *string) (*Config, error) {
	data, err := os.ReadFile(*file)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal([]byte(data), config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
