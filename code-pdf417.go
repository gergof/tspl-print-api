package main

type CodePdf417 struct {
	CodeBase
	X       int    `yaml:"x"`
	Y       int    `yaml:"y"`
	Width   int    `yaml:"width"`
	Height  int    `yaml:"height"`
	Content string `yaml:"content"`
}

func (c *CodePdf417) ToCommand(args map[string]string) (string, error) {
	renderedContent, err := fillTemplate(c.Content, args)
	if err != nil {
		return "", err
	}

	return TsplPdf417Command(c.X, c.Y, c.Width, c.Height, 0, renderedContent), nil
}
