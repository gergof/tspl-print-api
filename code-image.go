package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/url"
	"strings"
	"unicode"
)

type CodeImage struct {
	CodeBase
	X       int    `yaml:"x"`
	Y       int    `yaml:"y"`
	Content string `yaml:"content"`
}

func (c *CodeImage) ToCommand(args map[string]string) (string, error) {
	renderedContent, err := fillTemplate(c.Content, args)
	if err != nil {
		return "", err
	}

	decoded, err := decodeImageDataURL(renderedContent)
	if err != nil {
		return "", err
	}

	bitmap, width, height := imageToMonochromeBitmap(decoded)
	return TsplBitmapCommand(c.X, c.Y, width, height, 0, bitmap), nil
}

func decodeImageDataURL(dataURL string) (image.Image, error) {
	header, payload, found := strings.Cut(dataURL, ",")
	if !found || !strings.HasPrefix(strings.ToLower(header), "data:") {
		return nil, fmt.Errorf("image content must be a base64 data URL")
	}

	metadata := strings.Split(header[len("data:"):], ";")
	if len(metadata) == 0 || !strings.HasPrefix(strings.ToLower(metadata[0]), "image/") {
		return nil, fmt.Errorf("data URL must contain an image media type")
	}

	isBase64 := false
	for _, parameter := range metadata[1:] {
		if strings.EqualFold(parameter, "base64") {
			isBase64 = true
			break
		}
	}
	if !isBase64 {
		return nil, fmt.Errorf("image data URL must use base64 encoding")
	}

	payload, err := url.PathUnescape(payload)
	if err != nil {
		return nil, fmt.Errorf("invalid image data URL: %w", err)
	}
	payload = strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, payload)

	data, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		return nil, fmt.Errorf("invalid base64 image data: %w", err)
	}

	decoded, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("could not decode image: %w", err)
	}

	return decoded, nil
}

func imageToMonochromeBitmap(source image.Image) ([]byte, int, int) {
	bounds := source.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	bytesPerRow := (width + 7) / 8
	bitmap := make([]byte, bytesPerRow*height)
	grayscale := make([]float64, width*height)

	// The printer represents white pixels with set bits. Starting with an all
	// white bitmap also ensures unused padding bits do not print.
	for i := range bitmap {
		bitmap[i] = 0xff
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, a := source.At(bounds.Min.X+x, bounds.Min.Y+y).RGBA()

			// RGBA returns alpha-premultiplied channels. Composite transparent
			// pixels over white before calculating perceived luminance.
			luminance := (299*r+587*g+114*b)/1000 + (0xffff - a)
			grayscale[y*width+x] = float64(luminance) / 257
		}
	}

	// Floyd-Steinberg error diffusion preserves the appearance of gray tones
	// on a printer that can only place black dots or leave white space.
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			index := y*width + x
			oldPixel := grayscale[index]
			newPixel := 0.0
			if oldPixel >= 128 {
				newPixel = 255
			}

			if newPixel == 0 {
				bitmap[y*bytesPerRow+x/8] &^= 1 << (7 - uint(x%8))
			}

			error := oldPixel - newPixel
			if x+1 < width {
				grayscale[index+1] += error * 7 / 16
			}
			if y+1 < height {
				if x > 0 {
					grayscale[index+width-1] += error * 3 / 16
				}
				grayscale[index+width] += error * 5 / 16
				if x+1 < width {
					grayscale[index+width+1] += error / 16
				}
			}
		}
	}

	return bitmap, bytesPerRow, height
}
