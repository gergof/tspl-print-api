package main

import (
	"fmt"
)

type Code interface {
	ToCommand(args map[string]string) (string, error)
}

type CodeBase struct {
	Type string `yaml:"type"`
}

type CodeWrapper struct {
	Code
}

func (w *CodeWrapper) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var base CodeBase
	if err := unmarshal(&base); err != nil {
		return err
	}

	switch base.Type {
	case "text":
		var text CodeText
		if err := unmarshal(&text); err != nil {
			return err
		}
		w.Code = &text
	case "barcode":
		var barcode CodeBarcode
		if err := unmarshal(&barcode); err != nil {
			return err
		}
		w.Code = &barcode
	case "pdf417":
		var pdf417 CodePdf417
		if err := unmarshal(&pdf417); err != nil {
			return err
		}
		w.Code = &pdf417
	case "qr":
		var qr CodeQR
		if err := unmarshal(&qr); err != nil {
			return err
		}
		w.Code = &qr
	case "block":
		var block CodeBlock
		if err := unmarshal(&block); err != nil {
			return err
		}
		w.Code = &block
	case "datamatrix":
		var datamatrix CodeDatamatrix
		if err := unmarshal(&datamatrix); err != nil {
			return err
		}
		w.Code = &datamatrix
	default:
		return fmt.Errorf("unknown code type: %s", base.Type)
	}

	return nil
}
