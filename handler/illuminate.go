package handler

import (
	"github.com/luismesas/goPi/spi"
)

type Illuminate struct {
	Value int
	prev  int
}

func NewIlluminate(device *spi.SPIDevice, prev int) (*Illuminate, error) {
	raw, error := device.Send([3]byte{0x68, 0, 0})
	if error != nil {
		return nil, error
	}

	ret := &Illuminate{Value: ((int(raw[0]) << 8) + int(raw[1])) & 0x3FF, prev: prev}
	return ret, nil
}

func (self *Illuminate) On() bool {
	return (self.Value-self.prev > 20)
}
