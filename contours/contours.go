package contours

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"gocv.io/x/gocv"
)

func ReadFromFile(imageFile string) gocv.Mat {

	img := gocv.IMRead(imageFile, gocv.IMReadGrayScale)
	if img.Empty() {
		fmt.Print("File wasn't read. Failed.\n")

		return img
	}

	fmt.Print("File was read.\n")

	return img
}

func GetRectengleFromFile(img gocv.Mat) []image.Rectangle {
	harrcascade := "./contours/cascade/haarcascade_russian_plate_number.xml"

	classifier := gocv.NewCascadeClassifier()

	classifier.Load(harrcascade)
	defer classifier.Close()

	rects := classifier.DetectMultiScale(img)

	return rects
}

func SaveImage(img gocv.Mat) {
	imageFinal, err := img.ToImage()
	if err != nil {
		fmt.Print("Error ToImage\n")

		return
	}

	f, err := os.Create("outimage.png")
	if err != nil {
		fmt.Print("Error Create\n")

		return
	}

	defer f.Close()

	err = png.Encode(f, imageFinal)
	if err != nil {
		fmt.Print("Error Encoder\n")

		return
	}
}
