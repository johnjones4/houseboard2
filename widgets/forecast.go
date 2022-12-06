package widgets

import (
	"image"
	"main/core"
	"main/service"
	"math"

	"github.com/fogleman/gg"
)

type forecast struct {
	staticWidget
}

func NewForecast(row, col, rowspan, colspan int) core.Widget {
	return &forecast{
		staticWidget: staticWidget{
			row:     row,
			col:     col,
			rowspan: rowspan,
			colspan: colspan,
		},
	}
}

func (w *forecast) CanRender(info interface{}) bool {
	_, ok := info.(service.Weather)
	return ok
}

func (w *forecast) Draw(rect image.Rectangle, ctx *gg.Context, info interface{}) error {
	weather, ok := info.(service.Weather)
	if !ok {
		return core.ErrorUnsupportedType
	}
	var err error
	rect, err = w.staticWidget.Draw(rect, ctx, "Forecast")
	if err != nil {
		return err
	}

	periods := math.Min(4.0, float64(len(weather.Forecast)))
	periodWidth := (float64(rect.Dx()) - ((periods - 1) * padding)) / periods
	for i := 0; i < int(periods); i++ {
		x := rect.Min.X + ((int(periodWidth) + padding) * i)
		rect1 := image.Rect(x, rect.Min.Y, x+int(periodWidth), rect.Dy())
		drawWeather(rect1, ctx, weather.Forecast[i])
	}

	return nil
}
