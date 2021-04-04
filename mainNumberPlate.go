package main

import (
	"fmt"
	"recognition/contours"
	"strings"

	"github.com/otiai10/gosseract"
	"gocv.io/x/gocv"
)

func main() {
	img := contours.ReadFromFile("./test/test1.jpg")
	defer img.Close()

	if img.Empty() {

		return
	}

	window := gocv.NewWindow("Recognition Plate")
	defer window.Close()

	rects := contours.GetRectengleFromFile(img)

	if len(rects) == 0 {
		fmt.Print("Error\n")

		return
	}

	img2 := img.Region(rects[0])
	defer img2.Close()

	img3 := gocv.NewMat()
	defer img3.Close()

	for i := 0; i < img2.Rows(); i++ {
		for j := 0; j < img2.Cols(); j++ {
			if img2.GetUCharAt(i, j) > 92 {
				img2.SetUCharAt(i, j, 255)
			}
		}
	}

	img4 := img2.Clone()

	contours.SaveImage(img4)

	client := gosseract.NewClient()
	defer client.Close()

	client.SetImage("outimage.png")
	client.SetWhitelist("IAKBCEHMOPTX0123456789")
	client.SetPageSegMode(gosseract.PSM_SINGLE_BLOCK)
	boxs, _ := client.GetBoundingBoxes(gosseract.RIL_TEXTLINE)

	if len(boxs) == 0 {
		fmt.Println("Boxes is empry")
		return
	}

	text := boxs[0].Word

	text = strings.ReplaceAll(text, " ", "")

	fmt.Println(boxs[0].Box)
	fmt.Println(text)

	window.IMShow(img2)
	window.WaitKey(0)
}
