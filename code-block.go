package main

type CodeBlock struct {
	CodeBase
	X       int       `yaml:"x"`
	Y       int       `yaml:"y"`
	Width   int       `yaml:"width"`
	Height  int       `yaml:"height"`
	Font    string    `yaml:"font"`
	Space   int       `yaml:"space"`
	Align   TextAlign `yaml:"align"`
	Content string    `yaml:"content"`
}

func (c *CodeBlock) ToCommand(args map[string]string) (string, error) {
	renderedContent, err := fillTemplate(c.Content, args)
	if err != nil {
		return "", err
	}

	font := "1"
	if c.Font != "" {
		font = c.Font
	}

	space := 0
	if c.Space != 0 {
		space = c.Space
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

	return TsplBlockCommand(c.X, c.Y, c.Width, c.Height, font, 0, 1, 1, space, alignment, renderedContent), nil
}
