package main

import (
	"image"
	it8951 "main/IT8951"
	"os"

	"github.com/fogleman/gg"
	"github.com/johnjones4/houseboard2/core"
)

func main() {
	it8951.Init()

	width := it8951.Width()
	height := it8951.Height()

	ic := make(chan image.Image)
	go core.Run(ic, int(width), int(height), os.Getenv("CONFIG_FILE"))
	for img := range ic {
		gg.SavePNG("output.png", img)

		size := (img.Bounds().Max.X)*4/8 + (img.Bounds().Max.Y)*img.Bounds().Dx()
		buffer := make([]uint8, size)
		for y := 0; y < img.Bounds().Max.Y; y++ {
			for x := 0; x < img.Bounds().Max.X; x++ {
				r, b, g, _ := img.At(x, y).RGBA()
				gray := core.RGBtoGray(r, g, b)
				index := x*4/8 + y*img.Bounds().Dx()
				buffer[index] &= ^((0xF0) >> (7 - (x*4+3)%8))
				buffer[index] |= (gray & 0xF0) >> (7 - (x*4+3)%8)
			}
		}
		it8951.Draw(buffer)
	}
}
