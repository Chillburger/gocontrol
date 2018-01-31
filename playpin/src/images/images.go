package main

import (
  "fmt"
  "image"
  "image/color"
  "golang.org/x/tour/pic"
)

type Image struct {
  width int
  height int
  color uint8
}

func (i Image) ColorModel() color.Model {
  return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
  return image.Rect(0, 0, i.width, i.height)
}

func (i Image) At(x, y int) color.Color {
  return color.RGBA{i.color + uint8(x), i.color + uint8(y), 255, 255}
}

// the image package defines the Image interface
func main() {
  m := image.NewRGBA(image.Rect(0, 0, 100, 100))
  fmt.Println(m.Bounds())
  fmt.Println(m.At(0,0).RGBA())


  test := Image{100, 100, 100}
  pic.ShowImage(test)

}
