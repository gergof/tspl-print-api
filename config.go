package main

type Config struct {
	Endpoints map[string]Endpoint `yaml:"endpoints"`
}
