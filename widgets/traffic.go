package widgets

import (
	"image"
	"image/color"
	"main/core"
	"main/service"

	"github.com/fogleman/gg"
)

type traffic struct {
	staticWidget
}

func NewTraffic(row, col, rowspan, colspan int) core.Widget {
	return &traffic{
		staticWidget: staticWidget{
			row:     row,
			col:     col,
			rowspan: rowspan,
			colspan: colspan,
		},
	}
}

func (w *traffic) CanRender(info interface{}) bool {
	_, ok := info.(service.Traffic)
	return ok
}

func (w *traffic) Draw(rect image.Rectangle, ctx *gg.Context, info interface{}) error {
	traffic, ok := info.(service.Traffic)
	if !ok {
		return core.ErrorUnsupportedType
	}

	var err error
	rect, err = w.staticWidget.Draw(rect, ctx, "Traffic")
	if err != nil {
		return err
	}

	ctx.SetColor(color.Black)
	for i, dest := range traffic.Destinations {
		x := float64(rect.Min.X + padding)
		y := float64(rect.Min.Y + padding + ((fontNormalBold.Metrics().Height.Ceil() + padding) * i))

		ctx.SetFontFace(fontNormalBold)
		ctx.DrawStringAnchored(dest.Destination, x, y, 0, 1)

		ctx.SetFontFace(fontNormal)
		ctx.DrawStringAnchored(ageStrSeconds(dest.EstimatedDuration), float64(rect.Max.X-padding), y, 1, 1)
	}

	return nil
}
