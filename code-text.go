package main

import (
	"github.com/gergof/gotspl/gotspl"
)

type TextAlign string

const (
	TextAlignDefault TextAlign = "default"
	TextAlignLeft    TextAlign = "left"
	TextAlignCenter  TextAlign = "center"
	TextAlignRight   TextAlign = "right"
)

type CodeText struct {
	CodeBase
	X       int       `yaml:"x"`
	Y       int       `yaml:"y"`
	Font    string    `yaml:"font"`
	Align   TextAlign `yaml:"align"`
	Content string    `yaml:"content"`
}

func (c *CodeText) ToCommand(args map[string]string) (gotspl.TSPLCommand, error) {
	renderedContent, err := fillTemplate(c.Content, args)
	if err != nil {
		return nil, err
	}

	cmd := gotspl.Text()

	cmd = cmd.XCoordinate(c.X)
	cmd = cmd.YCoordinate(c.Y)
	cmd = cmd.Rotation(0)
	cmd = cmd.XMultiplier(1)
	cmd = cmd.YMultiplier(1)
	cmd = cmd.FontName("1")

	if c.Font != "" {
		cmd = cmd.FontName(c.Font)
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
