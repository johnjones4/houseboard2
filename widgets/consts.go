package widgets

import (
	"github.com/fogleman/gg"
	"golang.org/x/image/font"
)

func fontMust(path string, size float64) font.Face {
	f, err := gg.LoadFontFace(path, size)
	if err != nil {
		panic(err)
	}
	return f
}

var (
	fontTitle      = fontMust("./OPen_Sans/static/OpenSans/OpenSans-SemiBold.ttf", 20)
	fontBig        = fontMust("./OPen_Sans/static/OpenSans/OpenSans-SemiBold.ttf", 120)
	fontSmall      = fontMust("./OPen_Sans/static/OpenSans/OpenSans-Regular.ttf", 10)
	fontNormal     = fontMust("./OPen_Sans/static/OpenSans/OpenSans-Regular.ttf", 18)
	fontNormalBold = fontMust("./OPen_Sans/static/OpenSans/OpenSans-SemiBold.ttf", 18)
	fontNormalPlus = fontMust("./OPen_Sans/static/OpenSans/OpenSans-Regular.ttf", 22)
)

const (
	padding = 10
)

func fontRelativeBig(h float64) font.Face {
	return fontMust("./OPen_Sans/static/OpenSans/OpenSans-SemiBold.ttf", h*0.3)
}

func fontRelativeNormal(h float64) font.Face {
	return fontMust("./OPen_Sans/static/OpenSans/OpenSans-Regular.ttf", h*0.05)
}

func fontRelativeNormalBold(h float64) font.Face {
	return fontMust("./OPen_Sans/static/OpenSans/OpenSans-SemiBold.ttf", h*0.05)
}

func fontRelativeNormalPlus(h float64) font.Face {
	return fontMust("./OPen_Sans/static/OpenSans/OpenSans-Regular.ttf", h*0.06)
}
