package main

type PrinterLabel struct {
	Width  int32 `yaml:"width"`
	Height int32 `yaml:"height"`
	Gap    int32 `yaml:"gap"`
	Offset int32 `yaml:"offset"`
}

type Printer struct {
	Device string       `yaml:"device"`
	DPI    int32        `yaml:"dpi"`
	Label  PrinterLabel `yaml:"label"`
}
