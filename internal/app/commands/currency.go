package commands

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/galazat/cb_currency/internal/service/currency"

	"golang.org/x/net/html/charset"
)

var TodayCurrensies *currency.Currencies

func GetCurrencies(date string) *currency.Currencies {
	xmlBytes, err := getXML(GetUrl(date))
	if err != nil {
		log.Println(err)
	}

	currencies := currency.NewCurrencies()

	r := bytes.NewReader([]byte(xmlBytes))
	d := xml.NewDecoder(r)
	d.CharsetReader = charset.NewReaderLabel

	err = d.Decode(currencies)
	if err != nil {
		log.Println(err)
		return currencies
	}

	return currencies
}

func getXML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("Read body: %v", err)
	}

	return data, nil
}
