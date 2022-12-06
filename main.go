package main

import (
	"encoding/json"
	"main/core"
	"main/service"
	"main/widgets"
	"os"
	"strconv"

	"log"
)

func main() {
	configBytes, err := os.ReadFile(os.Getenv("CONFIG_FILE"))
	if err != nil {
		panic(err)
	}

	var config service.Configuration
	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		panic(err)
	}

	services := make([]core.Service, 0)
	for _, config := range config.Configurations() {
		if !config.Empty() {
			services = append(services, config.Service())
		}
	}

	outputChan := make(chan map[string]interface{})
	go service.RunServices(services, outputChan)

	w := []core.Widget{
		widgets.NewWeatherStation(0, 0, 3, 2),
		widgets.NewTraffic(3, 0, 2, 2),
		widgets.NewForecast(0, 2, 2, 4),
		widgets.NewCalendar(2, 2, 5, 4),
	}

	width, err := strconv.ParseInt(os.Getenv("WIDTH"), 10, 64)
	if err != nil {
		log.Panic(err)
	}

	height, err := strconv.ParseInt(os.Getenv("HEIGHT"), 10, 64)
	if err != nil {
		log.Panic(err)
	}

	for infos := range outputChan {
		ctx, err := widgets.Render(infos, w, int(width), int(height), 7, 6)
		if err != nil {
			log.Panic(err)
		}

		err = ctx.SavePNG("output.png")
		if err != nil {
			log.Panic(err)
		}

		os.Exit(0)
	}
}
