package shared

import (
	"context"
	"image"
	"time"

	"github.com/fogleman/gg"
)

type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type ServiceConfig interface {
	Empty() bool
	Service() Service
}

type Service interface {
	Name() string
	Info(c context.Context) (interface{}, error)
	NeedsRefresh() bool
}

type Widget interface {
	Row() int
	Col() int
	Rowspan() int
	Colspan() int
	CanRender(info interface{}) bool
	Draw(rect image.Rectangle, ctx *gg.Context, info interface{}) error
}

type Bounds struct {
}

type Weather interface {
	Temp() float64
	WindDirection() float64
	WindSpeed() float64
	Humidity() float64
	Press() float64
	Text() string
	Source() string
	Created() time.Time
	Label() string
}
