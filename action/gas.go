package action

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

func newClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: tr}
}

func RequestPostCurrentDate() error {
	gasURL := os.Getenv("LUGH_GAS_URL")

	client := newClient()
	response, err := client.PostForm(gasURL, url.Values{"turn_on_date": {time.Now().String()}})
	if err != nil {
		log.Fatalf("Response receive error: %s", err)
		return err
	}

	defer response.Body.Close()
	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Response read error: %s", err)
		return err
	}

	return nil
}
