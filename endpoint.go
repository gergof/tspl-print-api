package main

import (
	"github.com/gergof/gotspl/gotspl"
)

type Endpoint struct {
	Printer  Printer           `yaml:"printer"`
	Args     map[string]string `yaml:"args"`
	CodeList []CodeWrapper     `yaml:"code"`
}

func (e *Endpoint) RenderCodeList(args map[string]string) ([]byte, error) {
	label := gotspl.
		NewTSPLLabel().
		Cmd(
			gotspl.SizeCmd().
				LabelWidth(float64(e.Printer.Label.Width)).
				LabelLength(float64(e.Printer.Label.Height)),
		).
		Cmd(
			gotspl.GapCmd().
				LabelDistance(float64(e.Printer.Label.Gap)).
				LabelOffsetDistance(float64(e.Printer.Label.Offset)),
		).
		Cmd(
			gotspl.ClsCmd(),
		)

	for _, codeWrapper := range e.CodeList {
		cmd, err := codeWrapper.Code.ToCommand(args)
		if err != nil {
			return nil, err
		}

		label = label.Cmd(cmd)
	}

	label = label.Cmd(
		gotspl.PrintCmd().
			NumberLabels(1).
			NumberCopies(1),
	)

	return label.GetTSPLCode()
}

func (e *Endpoint) MapArgs(args map[string]string) map[string]string {
	result := make(map[string]string)

	for key := range e.Args {
		value, exists := args[e.Args[key]]
		if exists {
			result[key] = value
		}
	}

	return result
}
