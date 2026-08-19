package main

type CodeLine struct {
	CodeBase
	X      int `yaml:"x"`
	Y      int `yaml:"y"`
	Width  int `yaml:"width"`
	Height int `yaml:"height"`
}

func (c *CodeLine) ToCommand(args map[string]string) (string, error) {
	return TsplLineCommand(c.X, c.Y, c.Width, c.Height), nil
}
