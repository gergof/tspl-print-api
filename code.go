package main

import (
	"fmt"
	"github.com/gergof/gotspl/gotspl"
)

type Code interface {
	ToCommand(args map[string]string) (gotspl.TSPLCommand, error)
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
	default:
		return fmt.Errorf("unknown code type: %s", base.Type)
	}

	return nil
}
