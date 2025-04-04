package main

import (
	"strings"

	"github.com/tidwall/gjson"
)

type Endpoint struct {
	Printer  Printer           `yaml:"printer"`
	Args     map[string]string `yaml:"args"`
	CodeList []CodeWrapper     `yaml:"code"`
}

func (e *Endpoint) RenderCodeList(args map[string]string) (string, error) {
	label := make([]string, 0, 20)

	label = append(label, TsplSizeCommand(e.Printer.Label.Width, e.Printer.Label.Height))
	label = append(label, TsplGapCommand(e.Printer.Label.Gap, e.Printer.Label.Offset))
	label = append(label, TsplDirectionCommand(e.Printer.Direction == "inverted"))
	label = append(label, TsplClsCommand())

	for _, codeWrapper := range e.CodeList {
		cmd, err := codeWrapper.Code.ToCommand(args)
		if err != nil {
			return "", err
		}

		label = append(label, cmd)
	}

	label = append(label, TsplPrintCommand(1, 1))

	return strings.Join(label, "\n"), nil
}

func (e *Endpoint) GetArgsFromJson(json string) map[string]string {
	result := make(map[string]string)

	for key := range e.Args {
		result[key] = gjson.Get(json, e.Args[key]).String()
	}

	return result
}
