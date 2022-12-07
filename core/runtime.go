package core

import (
	"encoding/json"
	"image"
	"io"

	// it8951 "main/IT8951-ePaper"
	"os"

	"github.com/johnjones4/houseboard2/core/service"
	"github.com/johnjones4/houseboard2/core/shared"
	"github.com/johnjones4/houseboard2/core/widgets"

	"log"
)

func Run(imageChan chan image.Image, width int, height int, configPath string) {
	l := newLogger()
	log.SetOutput(&multiwriter{
		writers: []io.Writer{
			l,
			os.Stdout,
		},
	})

	configBytes, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	var config service.Configuration
	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		panic(err)
	}

	services := make([]shared.Service, 0)
	for _, config := range config.Configurations() {
		if !config.Empty() {
			services = append(services, config.Service())
		}
	}

	outputChan := make(chan map[string]interface{})
	go service.RunServices(services, outputChan)

	w := []shared.Widget{
		widgets.NewWeatherStation(0, 0, 3, 2),
		widgets.NewTraffic(3, 0, 2, 2),
		widgets.NewForecast(0, 2, 2, 4),
		widgets.NewCalendar(2, 2, 5, 4),
	}

	for infos := range outputChan {
		img, err := widgets.Render(infos, w, int(width), int(height), 7, 6)
		if err != nil {
			log.Panic(err)
		}
		imageChan <- img
	}
}
