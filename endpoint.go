package main

type Endpoint struct {
	Printer  Printer           `yaml:"printer"`
	Args     []string          `yaml:"args"`
	CodeList []CodeWrapper     `yaml:"code"`
	TestData map[string]string `yaml:"testData"`
}
