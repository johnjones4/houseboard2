package widgets

import (
	"image"
	"main/core"
	"main/service"

	"github.com/fogleman/gg"
)

type weatherStation struct {
	staticWidget
}

func NewWeatherStation(row, col, rowspan, colspan int) core.Widget {
	return &weatherStation{
		staticWidget: staticWidget{
			row:     row,
			col:     col,
			rowspan: rowspan,
			colspan: colspan,
		},
	}
}

func (w *weatherStation) CanRender(info interface{}) bool {
	_, ok := info.(service.WeatherStatonResponse)
	return ok
}

func (w *weatherStation) Draw(rect image.Rectangle, ctx *gg.Context, info interface{}) error {
	weather, ok := info.(service.WeatherStatonResponse)
	if !ok {
		return core.ErrorUnsupportedType
	}

	var err error
	rect, err = w.staticWidget.Draw(rect, ctx, "Weather Station")
	if err != nil {
		return err
	}

	drawWeather(rect, ctx, weather)

	return nil
}
