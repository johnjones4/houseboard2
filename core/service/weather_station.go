package service

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/johnjones4/houseboard2/core/shared"
)

type weatherStationConfiguration struct {
	Upstream string `json:"upstream"`
}

func (c weatherStationConfiguration) Empty() bool {
	return c.Upstream == ""
}

func (c weatherStationConfiguration) Service() shared.Service {
	return &weatherStation{
		configuration: c,
	}
}

type weatherStation struct {
	configuration weatherStationConfiguration
}

type WeatherStatonResponse struct {
	Timestamp        time.Time `json:"timestamp"`
	AvgWindSpeed     float64   `json:"anemometerAverage"`
	MinWindSpeed     float64   `json:"anemometerMin"`
	MaxWindSpeed     float64   `json:"anemometerMax"`
	Temperature      float64   `json:"temperature"`
	Gas              float64   `json:"gas"`
	RelativeHumidity float64   `json:"relativeHumidity"`
	Pressure         float64   `json:"pressure"`
	VaneDirection    float64   `json:"vaneDirection"`
}

type weatherStationResponseBody struct {
	Items []WeatherStatonResponse `json:"items"`
}

func (w WeatherStatonResponse) Temp() float64 {
	return w.Temperature
}

func (w WeatherStatonResponse) WindDirection() float64 {
	return w.VaneDirection
}

func (w WeatherStatonResponse) WindSpeed() float64 {
	return w.MaxWindSpeed
}

func (w WeatherStatonResponse) Humidity() float64 {
	return w.RelativeHumidity
}

func (w WeatherStatonResponse) Press() float64 {
	return w.Pressure
}

func (w WeatherStatonResponse) Source() string {
	return "Station"
}

func (w WeatherStatonResponse) Created() time.Time {
	return w.Timestamp
}

func (w WeatherStatonResponse) Text() string {
	return ""
}

func (w WeatherStatonResponse) Label() string {
	return ""
}

func (w *weatherStation) Name() string {
	return "weatherstation"
}

func (w *weatherStation) Info(c context.Context) (interface{}, error) {
	res, err := http.Get(w.configuration.Upstream)
	if err != nil {
		return WeatherStatonResponse{}, nil
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return WeatherStatonResponse{}, nil
	}

	var response weatherStationResponseBody
	err = json.Unmarshal(body, &response)
	if err != nil {
		return WeatherStatonResponse{}, nil
	}

	if len(response.Items) == 0 {
		return WeatherStatonResponse{}, shared.ErrorEmptyResponse
	}

	return response.Items[len(response.Items)-1], nil
}

func (w *weatherStation) NeedsRefresh() bool {
	return true
}
