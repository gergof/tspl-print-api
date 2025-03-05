package main

type PrinterLabel struct {
	Width  int `yaml:"width"`
	Height int `yaml:"height"`
	Gap    int `yaml:"gap"`
	Offset int `yaml:"offset"`
}

type Printer struct {
	Device string       `yaml:"device"`
	DPI    int          `yaml:"dpi"`
	Label  PrinterLabel `yaml:"label"`
}
