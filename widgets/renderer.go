package widgets

import (
	"image"
	"image/color"
	"main/core"

	"github.com/fogleman/gg"
)

func Render(infos map[string]interface{}, widgets []core.Widget, width, height, rows, cols int) (*gg.Context, error) {
	ctx := gg.NewContextForImage(image.NewGray(image.Rect(0, 0, width, height)))
	colWidth := width / cols
	rowWidth := height / rows
	ctx.SetColor(color.White)
	ctx.DrawRectangle(0, 0, float64(width), float64(height))
	ctx.Fill()
	for _, info := range infos {
		for _, widget := range widgets {
			if widget.CanRender(info) {
				x1 := widget.Col() * colWidth
				y1 := widget.Row() * rowWidth
				x2 := x1 + (widget.Colspan() * colWidth)
				y2 := y1 + (widget.Rowspan() * rowWidth)
				rect := image.Rect(x1, y1, x2, y2)
				err := widget.Draw(rect, ctx, info)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	return ctx, nil
}
