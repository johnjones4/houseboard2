package widgets

import (
	"image"
	"main/core"
	"main/service"

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
	// events, ok := info.([]service.Event)
	// if !ok {
	// 	return core.ErrorUnsupportedType
	// }

	var err error
	rect, err = w.staticWidget.Draw(rect, ctx, "Calendar")
	if err != nil {
		return err
	}

	// log.Println(events)

	return nil
}
