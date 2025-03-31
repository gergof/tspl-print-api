package main

import (
	"fmt"
	"strings"
)

func escapeContent(content string) string {
	return strings.ReplaceAll(content, "\"", "\\[\"]")
}

func TsplSizeCommand(width int, height int) string {
	return fmt.Sprintf("SIZE %d mm,%d mm", width, height)
}

func TsplGapCommand(distance int, offset int) string {
	if distance == 0 && offset == 0 {
		return "GAP 0,0"
	}

	return fmt.Sprintf("GAP %d mm,%d mm", distance, offset)
}

func TsplClsCommand() string {
	return "CLS"
}

func TsplDirectionCommand(inverted bool) string {
	if inverted {
		return "DIRECTION 1"
	} else {
		return "DIRECTION 0"
	}
}

func TsplPrintCommand(labels int, copies int) string {
	return fmt.Sprintf("PRINT %d,%d", labels, copies)
}

func TsplTextCommand(x int, y int, font string, rotation int, xmult int, ymult int, alignment int, content string) string {
	return fmt.Sprintf("TEXT %d,%d,\"%s\",%d,%d,%d,%d,\"%s\"", x, y, font, rotation, xmult, ymult, alignment, escapeContent(content))
}

func TsplBarcodeCommand(x int, y int, code string, height int, humanReadable int, rotation int, narrow int, wide int, alignment int, content string) string {
	return fmt.Sprintf("BARCODE %d,%d,\"%s\",%d,%d,%d,%d,%d,%d,\"%s\"", x, y, code, height, humanReadable, rotation, narrow, wide, alignment, escapeContent(content))
}

func TsplPdf417Command(x int, y int, width int, height int, rotate int, content string) string {
	return fmt.Sprintf("PDF417 %d,%d,%d,%d,%d,M1,\"%s\"", x, y, width, height, rotate, escapeContent(content))
}

func TsplQrCodeCommand(x int, y int, ecc string, cellWidth int, rotation int, content string) string {
	return fmt.Sprintf("QRCODE %d,%d,%s,%d,A,%d,\"%s\"", x, y, ecc, cellWidth, rotation, escapeContent(content))
}

func TsplBlockCommand(x int, y int, width int, height int, font string, rotation int, xmult int, ymult int, space int, align int, content string) string {
	return fmt.Sprintf("BLOCK %d,%d,%d,%d,\"%s\",%d,%d,%d,%d,%d,\"%s\"", x, y, width, height, font, rotation, xmult, ymult, space, align, escapeContent(content))
}

func TsplDatamatrixCommand(x int, y int, width int, height int, content string) string {
	return fmt.Sprintf("DMATRIX %d,%d,%d,%d,\"%s\"", x, y, width, height, escapeContent(content))
}
