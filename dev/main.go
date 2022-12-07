package main

import (
	"image"
	"log"
	"os"
	"strconv"

	"github.com/fogleman/gg"
	"github.com/johnjones4/houseboard2/core"
)

func main() {
	width, err := strconv.ParseInt(os.Getenv("WIDTH"), 10, 64)
	if err != nil {
		log.Panic(err)
	}

	height, err := strconv.ParseInt(os.Getenv("HEIGHT"), 10, 64)
	if err != nil {
		log.Panic(err)
	}

	ic := make(chan image.Image)
	go core.Run(ic, int(width), int(height), os.Getenv("CONFIG_FILE"))
	img := <-ic
	gg.SavePNG("output.png", img)
}
