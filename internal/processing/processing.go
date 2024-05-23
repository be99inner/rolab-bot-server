package processing

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"net/http"
	"os"
	"sync"
	"time"
)

var imageMutex sync.Mutex

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

// GenerateUniqueFilename generates a unique filename based on the current timestamp
func GenerateUniqueFilename() string {
	return time.Now().Format("20060102_150405.000") + ".jpg"
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

// ServeImage serves the saved image for preview in a web server
func ServeImage(w http.ResponseWriter, r *http.Request, images map[string]string) {
	imageMutex.Lock()
	defer imageMutex.Unlock()

	// Serve the most recently saved image
	if len(images) == 0 {
		http.Error(w, "No images found", http.StatusNotFound)
		return
	}

	var latestImage string
	for _, filename := range images {
		latestImage = filename
	}

	file, err := os.Open(latestImage)
	if err != nil {
		http.Error(w, "Image not found", http.StatusNotFound)
	}
	defer file.Close()

	w.Header().Set("Content-Type", "image/jpeg")
	http.ServeFile(w, r, latestImage)
}
