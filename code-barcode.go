package main

import (
	"github.com/mrgloba/gotspl/gotspl"
)

type HumanReadable string

const (
	HumanReadableNone   HumanReadable = "none"
	HumanReadableLeft   HumanReadable = "left"
	HumanReadableCenter HumanReadable = "center"
	HumanReadableRight  HumanReadable = "right"
)

type CodeBarcode struct {
	CodeBase
	X             int           `yaml:"x"`
	Y             int           `yaml:"y"`
	Height        int           `yaml:"height"`
	CodeType      string        `yaml:"codeType"`
	HumanReadable HumanReadable `yaml:"humanReadable"`
	Align         TextAlign     `yaml:"align"`
	Content       string        `yaml:"content"`
}

func (c *CodeBarcode) ToCommand(args map[string]string) (gotspl.TSPLCommand, error) {
	renderedContent, err := fillTemplate(c.Content, args)
	if err != nil {
		return nil, err
	}

	cmd := gotspl.BarcodeCmd()

	cmd.XCoordinate(c.X)
	cmd.YCoordinate(c.Y)
	cmd.Height(c.Height)
	cmd.CodeType(c.CodeType)

	if c.HumanReadable != "" {
		switch c.HumanReadable {
		case HumanReadableLeft:
			cmd.HumanReadable(1)
		case HumanReadableCenter:
			cmd.HumanReadable(2)
		case HumanReadableRight:
			cmd.HumanReadable(3)
		default:
			cmd.HumanReadable(0)
		}
	}

	if c.Align != "" {
		switch c.Align {
		case TextAlignLeft:
			cmd.Alignment(1)
		case TextAlignCenter:
			cmd.Alignment(2)
		case TextAlignRight:
			cmd.Alignment(3)
		default:
			cmd.Alignment(0)
		}
	}

	cmd.Content(renderedContent, true)

	return cmd, nil
}
