package widgets

import (
	"image"
	"image/color"

	"github.com/fogleman/gg"
)

type staticWidget struct {
	row     int
	col     int
	rowspan int
	colspan int
}

func (w *staticWidget) Row() int {
	return w.row
}

func (w *staticWidget) Col() int {
	return w.col
}

func (w *staticWidget) Rowspan() int {
	return w.rowspan
}

func (w *staticWidget) Colspan() int {
	return w.colspan
}

func (sw *staticWidget) Draw(rect image.Rectangle, ctx *gg.Context, title string) (image.Rectangle, error) {
	ctx.SetColor(color.Black)

	x := float64(rect.Min.X) + padding
	y := float64(rect.Min.Y) + padding
	w := float64(rect.Dx()) - (padding * 2)
	h := float64(rect.Dy()) - (padding * 2)
	ctx.DrawRectangle(x, y, w, h)
	ctx.SetLineWidth(1)
	ctx.Stroke()

	ctx.SetFontFace(fontTitle)
	textHeight := fontTitle.Metrics().Height.Ceil()
	y += padding + float64(textHeight)
	h -= padding + float64(textHeight)
	ctx.DrawString(title, x+padding, y)

	newRect := image.Rect(int(x), int(y), int(x+w), int(y+h))

	return newRect, nil
}
