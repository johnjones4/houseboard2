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
	}
}
