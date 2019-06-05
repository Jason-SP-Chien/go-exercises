package main

import "golang.org/x/tour/pic"
import "image"
import "image/color"

type Image struct{
	w, h int
}

func (i Image) ColorModel() color.Model {return color.RGBAModel}

func (i Image) Bounds() image.Rectangle {return image.Rect(0,0,i.w,i.h)}

func (i Image) At(x, y int) color.Color {return color.RGBA{uint8(x*x),uint8(y*y),uint8(x^y),uint8(y^x)}}

func main() {
	m := Image{1024,1024}
	pic.ShowImage(m)
}