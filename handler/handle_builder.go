package handler

import (
	"github.com/luismesas/goPi/spi"
	"log"
)

func RegisterHandlers(device *spi.SPIDevice) []Handler {
	return []Handler{NewIlluminate(device)}
}

func ListenCall(handlers []Handler) error {
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
