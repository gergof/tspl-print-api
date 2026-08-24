package main

type CodeQR struct {
	CodeBase
	X         int    `yaml:"x"`
	Y         int    `yaml:"y"`
	Ecc       string `yaml:"ecc"`
	CellWidth int    `yaml:"cellWidth"`
	Rotation  int    `yaml:"rotation"`
	Content   string `yaml:"content"`
}

func (c *CodeQR) ToCommand(args map[string]string) (string, error) {
	renderedContent, err := fillTemplate(c.Content, args)
	if err != nil {
		return "", err
	}

	return TsplQrCodeCommand(c.X, c.Y, c.Ecc, c.CellWidth, normalizeRotation(c.Rotation), renderedContent), nil
}
