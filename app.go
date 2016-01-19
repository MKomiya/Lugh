package main

import (
	"./handler"
	"github.com/luismesas/goPi/spi"
	"log"
	"time"
)

func main() {
	device := spi.NewSPIDevice(0, 0)
	ret := device.Open()
	if ret != nil {
		return
	}

	handlers := handler.RegisterHandlers(device)
	for {
		err := handler.ListenCall(handlers)
		if err != nil {
			log.Fatalf("Sensors event handler received error: %s", err)
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	device.Close()
}
