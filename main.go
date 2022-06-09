package main

import (
	// "image/png"
	"github.com/nfnt/resize"
	"image/jpeg"
	"log"
	"os"
)

const RESIZE_WIDTH = 200
const RESIZE_ASPECT_RATIO = 0

const TEST_FILE_PATH = "targets/afn-s.jpeg"
const RESIZED_TEST_FILE_PATH = "resized/afn-s.jpeg"

func main() {
	targetFile, error := os.Open(TEST_FILE_PATH)
	if error != nil {
		log.Fatal(error)
	}

	log.Print(targetFile)

	image, error := jpeg.Decode(targetFile)
	if error != nil {
		log.Fatal(error)
	}
	targetFile.Close()

	resizedImage := resize.Resize(RESIZE_WIDTH, RESIZE_ASPECT_RATIO, image, resize.Lanczos3)

	output, error := os.Create(RESIZED_TEST_FILE_PATH)
	if error != nil {
		log.Fatal(error)
	}
	defer output.Close()

	jpeg.Encode(output, resizedImage, nil)
}
