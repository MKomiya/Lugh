package main

import (
	"./handler"
	"github.com/luismesas/goPi/spi"
	"log"
	"time"
)

func registerHandlers(device *spi.SPIDevice) []handler.Handler {
	return []handler.Handler{handler.NewIlluminate(device)}
}

func listenCall(handlers []handler.Handler) error {
	for _, h := range handlers {
		if h.On() {
			err := h.Call()
			if err != nil {
				log.Fatalf("Failed action calling: %s", err)
				return err
			}
		}
	}
	return nil
}

func main() {
	device := spi.NewSPIDevice(0, 0)
	ret := device.Open()
	if ret != nil {
		return
	}

	handlers := registerHandlers(device)
	for {
		err := listenCall(handlers)
		if err != nil {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	device.Close()
}
