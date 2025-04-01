package main

import (
	"log"
	"os"
)

type PrinterLabel struct {
	Width  int `yaml:"width"`
	Height int `yaml:"height"`
	Gap    int `yaml:"gap"`
	Offset int `yaml:"offset"`
}

type Printer struct {
	Device    string       `yaml:"device"`
	Direction string       `yaml:"direction"`
	Label     PrinterLabel `yaml:"label"`
}

func (p *Printer) SendCommand(command []byte) error {
	file, err := os.OpenFile(p.Device, os.O_WRONLY|os.O_SYNC, 0666)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(command)
	if err != nil {
		return err
	}

	_, err = file.Write([]byte("\n\n\n\n\n"))

	err = file.Sync()
	if err != nil {
		// not fatal error, just log it
		log.Printf("Failed to force sync to printer: %v", err)
	}

	return err
}
