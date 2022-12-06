package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"main/core"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Weather struct {
	RadarURL string                              `json:"radarURL"`
	Forecast []NoaaWeatherForecastPeriod         `json:"forecast"`
	Alerts   []noaaWeatherAlertFeatureProperties `json:"alerts"`
}

type noaaWeatherPointProperties struct {
	ForecastURL  string `json:"forecast"`
	ForecastZone string `json:"forecastZone"`
	RadarStation string `json:"radarStation"`
}

type noaaWeatherPointResponse struct {
	Properties noaaWeatherPointProperties `json:"properties"`
}

type NoaaWeatherForecastPeriod struct {
	StartTime        time.Time `json:"startTime"`
	EndTime          time.Time `json:"endTime"`
	ShortForecast    string    `json:"shortForecast"`
	DetailedForecast string    `json:"detailedForecast"`
	Name             string    `json:"name"`
	Temperature      float64   `json:"temperature"`
	TemperatureUnit  string    `json:"temperatureUnit"`
	WindSpd          string    `json:"windSpeed"`
	WindDir          string    `json:"windDirection"`
	Icon             string    `json:"icon"`
	IsDaytime        bool      `json:"isDaytime"`
	generated        time.Time
}

func (w NoaaWeatherForecastPeriod) Temp() float64 {
	return w.Temperature
}

func (w NoaaWeatherForecastPeriod) WindDirection() float64 {
	switch w.WindDir {
	case "N":
		return 0
	case "NNE":
		return 22.5
	case "NE":
		return 45
	case "ENE":
		return 67.5
	case "E":
		return 90
	case "ESE":
		return 112.5
	case "SE":
		return 135
	case "SSE":
		return 157.5
	case "S":
		return 180
	case "SSW":
		return 202.5
	case "SW":
		return 225
	case "WSW":
		return 247.5
	case "W":
		return 270
	case "WNW":
		return 292.5
	case "NW":
		return 315
	case "NNW":
		return 337.5
	default:
		log.Panicf("Unexpected wind direction: %s", w.WindDir)
		return 0
	}
}

func (w NoaaWeatherForecastPeriod) WindSpeed() float64 {
	parts := strings.Split(w.WindSpd, " ")
	if len(parts) < 2 {
		return 0
	}
	speed, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		log.Panic(err)
	}
	return speed
}

func (w NoaaWeatherForecastPeriod) Humidity() float64 {
	return 0
}

func (w NoaaWeatherForecastPeriod) Press() float64 {
	return 0
}

func (w NoaaWeatherForecastPeriod) Source() string {
	return "NOAA"
}

func (w NoaaWeatherForecastPeriod) Created() time.Time {
	return w.generated
}

func (w NoaaWeatherForecastPeriod) Text() string {
	return w.ShortForecast
}

func (w NoaaWeatherForecastPeriod) Label() string {
	return w.Name
}

type noaaWeatherForecastProperties struct {
	Periods []NoaaWeatherForecastPeriod `json:"periods"`
	Updated time.Time                   `json:"updateTime"`
}

type noaaWeatherForecastResponse struct {
	Properties noaaWeatherForecastProperties `json:"properties"`
}

type noaaWeatherAlertFeatureProperties struct {
	ID            string   `json:"id"`
	AffectedZones []string `json:"affectedZones"`
	Headline      string   `json:"headline"`
}

type noaaWeatherAlertFeature struct {
	Properties noaaWeatherAlertFeatureProperties `json:"properties"`
}

type noaaWeatherAlertResponse struct {
	Features []noaaWeatherAlertFeature `json:"features"`
}

type noaaConfiguration struct {
	Location core.Coordinate
}

func (c noaaConfiguration) Empty() bool {
	return c.Location.Latitude == 0 && c.Location.Longitude == 0
}

func (c noaaConfiguration) Service() core.Service {
	return &noaa{c}
}

type noaa struct {
	configuration noaaConfiguration
}

func (f *noaa) Name() string {
	return "noaa"
}

func (f *noaa) Info(c context.Context) (interface{}, error) {
	weather, err := f.predictWeather(f.configuration.Location)
	if err != nil {
		return nil, err
	}
	return weather, nil
}

func (f *noaa) NeedsRefresh() bool {
	return true
}

func (f *noaa) predictWeather(coord core.Coordinate) (Weather, error) {
	point, err := makeWeatherAPIPointRequest(coord)
	if err != nil {
		return Weather{}, err
	}

	radarURL := fmt.Sprintf("https://radar.weather.gov/ridge/lite/%s_loop.gif?v=%d", point.RadarStation, time.Now().Unix())

	forecast, err := makeWeatherAPIForecastCall(point)
	if err != nil {
		return Weather{}, err
	}

	alerts, err := makeWeatherAPIAlertCall(point)
	if err != nil {
		return Weather{}, err
	}

	return Weather{
		RadarURL: radarURL,
		Forecast: forecast,
		Alerts:   alerts,
	}, nil
}

func makeWeatherAPIPointRequest(coord core.Coordinate) (noaaWeatherPointProperties, error) {
	httpResponse, err := http.Get(fmt.Sprintf("https://api.weather.gov/points/%f,%f", coord.Latitude, coord.Longitude))
	if err != nil {
		return noaaWeatherPointProperties{}, err
	}

	responseBytes, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return noaaWeatherPointProperties{}, err
	}

	var pointResponse noaaWeatherPointResponse
	err = json.Unmarshal(responseBytes, &pointResponse)

	return pointResponse.Properties, err
}

func makeWeatherAPIForecastCall(point noaaWeatherPointProperties) ([]NoaaWeatherForecastPeriod, error) {
	httpResponse, err := http.Get(point.ForecastURL)
	if err != nil {
		return nil, nil
	}

	responseBytes, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, nil
	}

	var response noaaWeatherForecastResponse
	err = json.Unmarshal(responseBytes, &response)
	if err != nil {
		return nil, nil
	}

	if len(response.Properties.Periods) == 0 {
		return nil, errors.New("no forecast data returned")
	}

	for i := range response.Properties.Periods {
		response.Properties.Periods[i].generated = response.Properties.Updated
	}

	return response.Properties.Periods, nil
}

func makeWeatherAPIAlertCall(point noaaWeatherPointProperties) ([]noaaWeatherAlertFeatureProperties, error) {
	zoneId := strings.Replace(point.ForecastZone, "https://api.weather.gov/zones/forecast/", "", 1)

	httpResponse, err := http.Get(fmt.Sprintf("https://api.weather.gov/alerts/active/zone/%s", zoneId))
	if err != nil {
		return nil, err
	}

	responseBytes, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	var response noaaWeatherAlertResponse
	err = json.Unmarshal(responseBytes, &response)
	if err != nil {
		return nil, err
	}

	featureProps := make([]noaaWeatherAlertFeatureProperties, 0)
	for _, feature := range response.Features {
		for _, zone := range feature.Properties.AffectedZones {
			if zone == point.ForecastZone {
				featureProps = append(featureProps, feature.Properties)
			}
		}
	}
	return featureProps, nil
}
