package main

type Endpoint struct {
	Printer  Printer           `yaml:"printer"`
	Args     map[string]string `yaml:"args"`
	CodeList []CodeWrapper     `yaml:"code"`
}
