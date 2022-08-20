package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
}

func (m Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (m Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 255, 255)
}

func (m Image) At(x, y int) color.Color {
	var v uint8 = 0
	if (x+y)%2 == 0 {
		v = 255
	}
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
