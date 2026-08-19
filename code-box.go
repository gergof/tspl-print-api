package main

type CodeBox struct {
	CodeBase
	X         int `yaml:"x"`
	Y         int `yaml:"y"`
	Width     int `yaml:"width"`
	Height    int `yaml:"height"`
	Thickness int `yaml:"thickness"`
}

func (c *CodeBox) ToCommand(args map[string]string) (string, error) {
	thickness := c.Thickness
	if thickness == 0 {
		thickness = 1
	}

	return TsplBoxCommand(c.X, c.Y, c.Width, c.Height, thickness), nil
}
