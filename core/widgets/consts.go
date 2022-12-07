package widgets

import (
	_ "embed"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

//go:embed Open_Sans/static/OpenSans/OpenSans-SemiBold.ttf
var semibold []byte

//go:embed Open_Sans/static/OpenSans/OpenSans-Regular.ttf
var regular []byte

func fontMust(bytes []byte, size float64) font.Face {
	f, err := truetype.Parse(bytes)
	if err != nil {
		panic(err)
	}
	return truetype.NewFace(f, &truetype.Options{
		Size: size,
	})
}

var (
	fontTitle      = fontMust(semibold, 20)
	fontSmall      = fontMust(regular, 10)
	fontNormal     = fontMust(regular, 18)
	fontNormalBold = fontMust(semibold, 18)
)

const (
	padding = 10
)

func fontRelativeBig(h float64) font.Face {
	return fontMust(semibold, h*0.3)
}

func fontRelativeNormal(h float64) font.Face {
	return fontMust(regular, h*0.05)
}

func fontRelativeNormalBold(h float64) font.Face {
	return fontMust(semibold, h*0.05)
}

func fontRelativeNormalPlus(h float64) font.Face {
	return fontMust(regular, h*0.06)
}
