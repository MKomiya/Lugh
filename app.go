package main

import (
	"./handler"
	"github.com/luismesas/goPi/spi"
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
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	device.Close()
}
