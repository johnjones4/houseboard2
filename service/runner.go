package service

import (
	"context"
	"log"
	"main/core"
	"time"
)

const defaultWait = time.Second * 5

func RunServices(services []core.Service, outputChan chan map[string]interface{}) {
	output := make(map[string]interface{})
	for {
		for _, service := range services {
			name := service.Name()
			if _, ok := output[name]; !ok || service.NeedsRefresh() {
				var err error
				waitDelay := defaultWait
				for err != nil || waitDelay == defaultWait {
					log.Printf("Updating %s", name)
					waitDelay = waitDelay * 2
					info, err := service.Info(context.Background())
					if err != nil {
						log.Printf("Error: \"%s\". Sleeping %s", err.Error(), waitDelay.String())
						time.Sleep(waitDelay)
						if waitDelay > time.Second*30 {
							break
						}
					} else {
						output[name] = info
					}
				}
			}
		}
		outputChan <- output
		time.Sleep(time.Minute * 5)
	}
}
