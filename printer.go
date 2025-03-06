package main

import "os"

type PrinterLabel struct {
	Width  int `yaml:"width"`
	Height int `yaml:"height"`
	Gap    int `yaml:"gap"`
	Offset int `yaml:"offset"`
}

type Printer struct {
	Device string       `yaml:"device"`
	Label  PrinterLabel `yaml:"label"`
}

func (p *Printer) SendCommand(command []byte) error {
	file, err := os.OpenFile(p.Device, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(command)

	return err
}
