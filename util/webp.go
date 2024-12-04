package util

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"strings"

	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
	"github.com/konotorii/go-consola"
)

func EncodeWebP(filePath string) (*string, error) {
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	ext := path.Ext(filePath)

	if ext != ".jpg" && ext != ".png" && ext != ".jpeg" {
		return nil, errors.New("not an image")
	}

	file, err := os.Open(filePath)
	if err != nil {
		consola.Error(err)
		return nil, err
	}

	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
	if err != nil {
		consola.Error(err)
		return nil, err
	}

	fileName := strings.Replace(filePath, ext, ".webp", -1)

	output, err := os.Create(fileName)
	if err != nil {
		consola.Error(err)
		return nil, err
	}
	defer output.Close()

	var img image.Image

	if ext == ".jpg" || ext == ".jpeg" {
		img, err = jpeg.Decode(file)
		if err != nil {
			consola.Error(err)
			return nil, err
		}
	}
	if ext == ".png" {
		img, err = png.Decode(file)
		if err != nil {
			consola.Error(err)
			return nil, err
		}
	}

	if err := webp.Encode(output, img, options); err != nil {
		consola.Error(err)
		return nil, err
	}

	return &fileName, nil
}

func DecodeWebP(filePath string) {

}

func WebpExists(filePath string) bool {
	ext := path.Ext(filePath)
	fileName := strings.Replace(filePath, ext, ".webp", -1)

	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}