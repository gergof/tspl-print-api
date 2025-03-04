package main

import (
	"github.com/mrgloba/gotspl"
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
	X       int32     `yaml:"x"`
	Y       int32     `yaml:"y"`
	Font    string    `yaml:"font"`
	Align   TextAlign `yaml:"align"`
	Content string    `yaml:"content"`
}

func (c *CodeText) ToCommand(args map[string]string) (error, gotspl.TSPLCommand) {
	err, renderedContent := fillTemplate(c.Content, args)
	if err != nil {
		return err, nil
	}

	cmd := gotspl.TextCmd()

	cmd.XCoordinate(c.X)
	cmd.YCoordinate(c.Y)

	if c.Font != "" {
		cmd.FontName(c.Font)
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

	cmd.Content(renderedContent)

	return nil, cmd
}
