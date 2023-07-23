package util

import (
	"image"

	"github.com/kbinani/screenshot"
)

func ReadScreen() *image.RGBA {
	n := screenshot.NumActiveDisplays()
	if n <= 0 {
		panic("Display not found")
	}

	bounds := screenshot.GetDisplayBounds(0)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}
	return img
	// Drawing IMG Gray
	/*b := img.Bounds()
	imgSet := image.NewRGBA(b)
	for y := 0; y < b.Max.Y; y++ {
		for x := 0; x < b.Max.X; x++ {
			oldPixel := img.At(x, y)
			pixel := color.GrayModel.Convert(oldPixel)
			imgSet.Set(x, y, pixel)
		}
	}

	return imgSet
	*/
}
