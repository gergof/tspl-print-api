package main

type CodeDatamatrix struct {
	CodeBase
	X       int    `yaml:"x"`
	Y       int    `yaml:"y"`
	Width   int    `yaml:"width"`
	Height  int    `yaml:"height"`
	Content string `yaml:"content"`
}

func (c *CodeDatamatrix) ToCommand(args map[string]string) (string, error) {
	renderedContent, err := fillTemplate(c.Content, args)
	if err != nil {
		return "", err
	}

	return TsplDatamatrixCommand(c.X, c.Y, c.Width, c.Height, renderedContent), nil
}
