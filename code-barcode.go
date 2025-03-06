package main

import (
	"github.com/gergof/gotspl/gotspl"
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

	cmd = cmd.XCoordinate(c.X)
	cmd = cmd.YCoordinate(c.Y)
	cmd = cmd.Height(c.Height)
	cmd = cmd.CodeType(c.CodeType)
	cmd = cmd.Rotation(0)
	cmd = cmd.Narrow(2)
	cmd = cmd.Wide(2)

	if c.HumanReadable != "" {
		switch c.HumanReadable {
		case HumanReadableLeft:
			cmd = cmd.HumanReadable(1)
		case HumanReadableCenter:
			cmd = cmd.HumanReadable(2)
		case HumanReadableRight:
			cmd = cmd.HumanReadable(3)
		default:
			cmd = cmd.HumanReadable(0)
		}
	}

	if c.Align != "" {
		switch c.Align {
		case TextAlignLeft:
			cmd = cmd.Alignment(1)
		case TextAlignCenter:
			cmd = cmd.Alignment(2)
		case TextAlignRight:
			cmd = cmd.Alignment(3)
		default:
			cmd = cmd.Alignment(0)
		}
	}

	cmd = cmd.Content(renderedContent, true)

	return cmd, nil
}
