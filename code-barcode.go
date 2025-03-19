package main

type HumanReadable string

const (
	HumanReadableNone   HumanReadable = "none"
	HumanReadableLeft   HumanReadable = "left"
	HumanReadableCenter HumanReadable = "center"
	HumanReadableRight  HumanReadable = "right"
)

type CodeBarcode struct {
	CodeBase
	X             int           `yaml:"x"`
	Y             int           `yaml:"y"`
	Height        int           `yaml:"height"`
	CodeType      string        `yaml:"codeType"`
	HumanReadable HumanReadable `yaml:"humanReadable"`
	Align         TextAlign     `yaml:"align"`
	Content       string        `yaml:"content"`
}

func (c *CodeBarcode) ToCommand(args map[string]string) (string, error) {
	renderedContent, err := fillTemplate(c.Content, args)
	if err != nil {
		return "", err
	}

	humanReadable := 0
	if c.HumanReadable != "" {
		switch c.HumanReadable {
		case HumanReadableLeft:
			humanReadable = 1
		case HumanReadableCenter:
			humanReadable = 2
		case HumanReadableRight:
			humanReadable = 3
		default:
			humanReadable = 0
		}
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

	return TsplBarcodeCommand(c.X, c.Y, c.CodeType, c.Height, humanReadable, 0, 2, 2, alignment, renderedContent), nil
}
