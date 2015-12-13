package main

import (
	"./action"
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

	prev := 0
	for {
		illuminate, error := handler.NewIlluminate(device, prev)
		if error != nil {
			log.Fatal("Failed create illuminate")
			break
		}

		if illuminate.On() {
			action.DoAction()
		}
		prev = illuminate.Value

		time.Sleep(500 * time.Millisecond)
	}
	device.Close()
}
