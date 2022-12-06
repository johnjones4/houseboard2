package widgets

import (
	"fmt"
	"image"
	"image/color"
	"main/core"
	"math"
	"time"

	"github.com/fogleman/gg"
)

func drawWeather(rect image.Rectangle, ctx *gg.Context, weather core.Weather) {
	ctx.SetColor(color.Black)

	_fontNormlBold := fontRelativeNormalBold(float64(rect.Dy()))

	y := rect.Min.Y
	if weather.Label() != "" {
		middleX := rect.Min.X + (rect.Dx() / 2)
		ctx.SetFontFace(_fontNormlBold)
		ctx.DrawStringAnchored(weather.Label(), float64(middleX), float64(y+(padding*2)), 0.5, 1)
		y += _fontNormlBold.Metrics().Height.Ceil() + (padding * 2)
	}

	_fontBig := fontRelativeBig(float64(rect.Dy()))
	ctx.SetFontFace(_fontBig)
	ctx.DrawStringAnchored(
		fmt.Sprintf("%0.1f°", weather.Temp()),
		float64(rect.Min.X+(rect.Dx()/2)),
		float64(y+_fontBig.Metrics().Height.Ceil()),
		0.5,
		0,
	)

	leftSideWidth := float64(rect.Dx()-(padding*3)) * 0.45
	radius := int(leftSideWidth/2) - padding
	circleX := rect.Min.X + padding + int(leftSideWidth/2)
	circleY := y + _fontBig.Metrics().Height.Ceil() + (padding * 3) + radius
	ctx.SetLineWidth(float64(rect.Dy()) * 0.0142857142857)
	ctx.DrawCircle(float64(circleX), float64(circleY), float64(radius))
	ctx.Stroke()

	_fontNormalPlus := fontRelativeNormalPlus(float64(rect.Dy()))
	ctx.SetFontFace(_fontNormalPlus)
	ctx.DrawStringAnchored(
		fmt.Sprintf("%0.1f mph", weather.WindSpeed()),
		float64(circleX),
		float64(circleY),
		0.5,
		0.5,
	)

	ctx.Push()
	ctx.RotateAbout((weather.WindDirection()+180)*(math.Pi/180.0), float64(circleX), float64(circleY))
	var edge float64 = float64(rect.Dy()) * 0.0857142857143
	b := math.Sqrt(math.Pow(edge, 2) - math.Pow(edge/2, 2))
	shape(
		ctx,
		[]image.Point{
			image.Pt(int(circleX), circleY-radius-int(b/2)),
			image.Pt(int(circleX)-int(edge/2), int(circleY)-radius+int(b/2)),
			image.Pt(int(circleX)+int(edge/2), int(circleY)-radius+int(b/2)),
		},
	)
	ctx.Fill()
	ctx.Pop()

	type keyVal struct {
		key   string
		value string
	}
	keyVals := make([]keyVal, 0)
	if weather.Humidity() > 0 {
		keyVals = append(keyVals, keyVal{
			key:   "Humidity",
			value: fmt.Sprintf("%0.1f %%", weather.Humidity()),
		})
	}
	if weather.Press() > 0 {
		keyVals = append(keyVals, keyVal{
			key:   "Pressure",
			value: fmt.Sprintf("%0.1f inHg", weather.Press()),
		})
	}
	if !weather.Created().IsZero() {
		keyVals = append(keyVals, keyVal{
			key:   "Age",
			value: ageStr(time.Since(weather.Created())),
		})
	}
	keyVals = append(keyVals, keyVal{
		key:   "Source",
		value: weather.Source(),
	})

	textX := rect.Min.X + padding + int(leftSideWidth) + padding
	lines := len(keyVals)
	extraText := weather.Text()
	extraTextLines := []string{}
	if extraText != "" {
		width := float64(rect.Max.X - padding - textX)
		extraTextLines = ctx.WordWrap(extraText, width)
		lines += len(extraTextLines)
	}

	_fontNorml := fontRelativeNormal(float64(rect.Dy()))

	textY := circleY - (lines / 2 * _fontNorml.Metrics().Height.Ceil())

	for _, line := range extraTextLines {
		ctx.SetFontFace(_fontNorml)
		ctx.DrawStringAnchored(line, float64(textX), float64(textY), 0, 1)
		textY += _fontNorml.Metrics().Height.Ceil()
	}

	for _, kv := range keyVals {
		keyStr := kv.key + ":"

		ctx.SetFontFace(_fontNormlBold)
		keyWidth, _ := ctx.MeasureString(keyStr)
		ctx.DrawStringAnchored(keyStr, float64(textX), float64(textY), 0, 1)

		ctx.SetFontFace(_fontNorml)
		ctx.DrawStringAnchored(fmt.Sprint(kv.value), float64(textX)+(padding/2)+keyWidth, float64(textY), 0, 1)

		textY += _fontNorml.Metrics().Height.Ceil()
	}
}

func shape(ctx *gg.Context, points []image.Point) {
	ctx.MoveTo(float64(points[0].X), float64(points[0].Y))
	for i := 1; i <= len(points); i++ {
		ctx.LineTo(float64(points[i%len(points)].X), float64(points[i%len(points)].Y))
	}
}
