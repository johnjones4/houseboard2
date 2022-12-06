package widgets

import (
	"fmt"
	"image"
	"image/color"
	"main/core"
	"main/service"
	"time"

	"github.com/fogleman/gg"
)

type calendar struct {
	staticWidget
}

func NewCalendar(row, col, rowspan, colspan int) core.Widget {
	return &calendar{
		staticWidget: staticWidget{
			row:     row,
			col:     col,
			rowspan: rowspan,
			colspan: colspan,
		},
	}
}

func (w *calendar) CanRender(info interface{}) bool {
	_, ok := info.([]service.Event)
	return ok
}

func (w *calendar) Draw(rect image.Rectangle, ctx *gg.Context, info interface{}) error {
	events, ok := info.([]service.Event)
	if !ok {
		return core.ErrorUnsupportedType
	}

	var err error
	rect, err = w.staticWidget.Draw(rect, ctx, "Calendar")
	if err != nil {
		return err
	}

	cols := 7
	rows := 4
	cellWidth := (rect.Dx() - (padding * 2)) / cols
	cellHeight := (rect.Dy() - (padding * 2)) / rows
	day := time.Now()
	startCol := int(day.Weekday())
	ctx.SetColor(color.Black)
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if (row == 0 && col >= startCol) || row > 0 {
				x := rect.Min.X + padding + (cellWidth * col)
				y := rect.Min.Y + padding + (cellHeight * row)
				rect1 := image.Rect(
					x,
					y,
					x+cellWidth,
					y+cellHeight,
				)
				drawDay(ctx, eventsForDay(day, events), day, rect1, col == cols-1, row == rows-1)
				day = day.Add(time.Hour * 24)
			}
		}
	}

	return nil
}

func drawDay(ctx *gg.Context, events []service.Event, day time.Time, rect image.Rectangle, endCol bool, endRow bool) {
	ctx.SetLineWidth(1)
	ctx.DrawLine(float64(rect.Min.X), float64(rect.Min.Y), float64(rect.Max.X), float64(rect.Min.Y))
	ctx.Stroke()
	ctx.DrawLine(float64(rect.Min.X), float64(rect.Min.Y), float64(rect.Min.X), float64(rect.Max.Y))
	ctx.Stroke()
	if endCol {
		ctx.DrawLine(float64(rect.Max.X), float64(rect.Min.Y), float64(rect.Max.X), float64(rect.Max.Y))
		ctx.Stroke()
	}
	if endRow {
		ctx.DrawLine(float64(rect.Min.X), float64(rect.Max.Y), float64(rect.Max.X), float64(rect.Max.Y))
		ctx.Stroke()
	}
	ctx.SetFontFace(fontNormalBold)
	ctx.DrawStringAnchored(fmt.Sprint(day.Day()), float64(rect.Max.X-padding), float64(rect.Min.Y+padding), 1, 0.75)

	ctx.SetFontFace(fontSmall)
	textX := rect.Min.X + padding
	textBaseY := rect.Min.Y + fontSmall.Metrics().Height.Ceil() + (padding * 2)
	line := 0
	for _, event := range events {
		lines := ctx.WordWrap("• "+event.Title, float64(rect.Dx()-(padding*2)))
		for _, l := range lines {
			textY := textBaseY + (line * fontSmall.Metrics().Height.Ceil())
			ctx.DrawStringAnchored(l, float64(textX), float64(textY), 0, 1)
			line++
		}
	}
}

func eventsForDay(day time.Time, events []service.Event) []service.Event {
	start := time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, day.Location())
	end := time.Date(day.Year(), day.Month(), day.Day(), 23, 59, 59, 0, day.Location())
	filtered := make([]service.Event, 0)
	for _, event := range events {
		if event.End.After(start) && event.Start.Before(end) {
			filtered = append(filtered, event)
		}
	}
	return filtered
}
