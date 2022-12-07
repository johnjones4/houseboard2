package service

import "github.com/johnjones4/houseboard2/core/shared"

type Configuration struct {
	ICal           iCalConfiguration           `json:"ical"`
	NOAA           noaaConfiguration           `json:"noaa"`
	WeatherStation weatherStationConfiguration `json:"weatherStation"`
	Traffic        trafficConfiguration        `json:"traffic"`
}

func (c Configuration) Configurations() []shared.ServiceConfig {
	return []shared.ServiceConfig{
		c.ICal,
		c.NOAA,
		c.WeatherStation,
		c.Traffic,
	}
}
