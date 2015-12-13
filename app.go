package main

import (
	"crypto/tls"
	"github.com/luismesas/goPi/spi"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

type Illuminate struct {
	value int
	prev  int
}

func NewIlluminate(device *spi.SPIDevice, prev int) (*Illuminate, error) {
	raw, error := device.Send([3]byte{0x68, 0, 0})
	if error != nil {
		return nil, error
	}

	ret := &Illuminate{value: ((int(raw[0]) << 8) + int(raw[1])) & 0x3FF, prev: prev}
	return ret, nil
}

func (self *Illuminate) IsIlluminated() bool {
	return (self.value-self.prev > 20)
}

func newClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: tr}
}

func RequestWelcomeBackDate() {
	gas_url := os.Getenv("LUGH_GAS_URL")

	client := newClient()
	response, err := client.PostForm(gas_url, url.Values{"turn_on_date": {time.Now().String()}})
	if err != nil {
		log.Fatalf("Response receive error: %s", err)
		return
	}

	defer response.Body.Close()
	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Response read error: %s", err)
		return
	}
}

func main() {
	device := spi.NewSPIDevice(0, 0)
	ret := device.Open()
	if ret != nil {
		return
	}

	prev := 0
	for {
		illuminate, error := NewIlluminate(device, prev)
		if error != nil {
			log.Fatal("Failed create illuminate")
			break
		}

		if illuminate.IsIlluminated() {
			RequestWelcomeBackDate()
		}
		prev = illuminate.value

		time.Sleep(1)
	}
	device.Close()
}
