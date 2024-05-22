package processing

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"os"
)

// DecodeImage decodes a base64 encoded image string
func DecodeImage(encodedStr string) (image.Image, error) {
	imgData, err := base64.StdEncoding.DecodeString(encodedStr)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(bytes.NewReader(imgData))
	if err != nil {
		return nil, err
	}
	return img, nil
}

// SaveImage saves an image to a file.
func SaveImage(filename string, img image.Image) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return jpeg.Encode(file, img, nil)
}
