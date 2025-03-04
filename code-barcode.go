package main

import (
	"github.com/mrgloba/gotspl"
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
	X             int32         `yaml:"x"`
	Y             int32         `yaml:"y"`
	Height        int32         `yaml:"height"`
	CodeType      string        `yaml:"codeType"`
	HumanReadable HumanReadable `yaml:"humanReadable"`
	Align         TextAlignt    `yaml:"align"`
	Content       string        `yaml:"content"`
}

func (c *CodeText) ToCommand(args map[string]string) (error, gotspl.TSPLCommand) {
	err, renderedContent := fillTemplate(c.Content, args)
	if err != nil {
		return err, nil
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
		case HumanReadableLeft:
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
		case TextAlignLeft:
			cmd.Alignment(3)
		default:
			cmd.Alignment(0)
		}
	}

	cmd.Content(renderedContent)

	return nil, cmd
}
