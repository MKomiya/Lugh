package handler

import (
	"./../action"
	"github.com/luismesas/goPi/spi"
)

// Illuminate has a value which is an illumination params from spi device
type Illuminate struct {
	device *spi.SPIDevice
	value  int
	prev   int
}

// NewIlluminate instantiates to Illuminate. It requires spi device instance
func NewIlluminate(device *spi.SPIDevice) *Illuminate {
	return &Illuminate{device: device, value: 0, prev: 0}
}

// On function gets latest illumination params and check if it gets a light
func (illuminate *Illuminate) On() bool {
	raw, error := illuminate.device.Send([3]byte{0x68, 0, 0})
	if error != nil {
		return false
	}

	illuminate.prev = illuminate.value
	illuminate.value = ((int(raw[0]) << 8) + int(raw[1])) & 0x3FF
	return (illuminate.value-illuminate.prev > 20)
}

// Call function is doing to request for posting current date to Google SpreadSheet
func (*Illuminate) Call() error {
	return action.RequestPostCurrentDate()
}
