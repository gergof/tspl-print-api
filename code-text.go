package main

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

func (c *CodeText) ToCommand(args map[string]string) (string, error) {
	renderedContent, err := fillTemplate(c.Content, args)
	if err != nil {
		return "", err
	}

	font := "1"
	if c.Font != "" {
		font = c.Font
	}

	alignment := 0
	if c.Align != "" {
		switch c.Align {
		case TextAlignLeft:
			alignment = 1
		case TextAlignCenter:
			alignment = 2
		case TextAlignRight:
			alignment = 3
		default:
			alignment = 0
		}
	}

	return TsplTextCommand(c.X, c.Y, font, 0, 1, 1, alignment, renderedContent), nil
}
