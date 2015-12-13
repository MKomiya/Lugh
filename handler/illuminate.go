package handler

import (
	"./../action"
	"github.com/luismesas/goPi/spi"
)

type Illuminate struct {
	device *spi.SPIDevice
	value  int
	prev   int
}

func NewIlluminate(device *spi.SPIDevice) *Illuminate {
	return &Illuminate{device: device, value: 0, prev: 0}
}

func (self *Illuminate) On() bool {
	raw, error := self.device.Send([3]byte{0x68, 0, 0})
	if error != nil {
		return false
	}

	self.prev = self.value
	self.value = ((int(raw[0]) << 8) + int(raw[1])) & 0x3FF
	return (self.value-self.prev > 20)
}

func (*Illuminate) Call() error {
	return action.RequestPostCurrentDate()
}
