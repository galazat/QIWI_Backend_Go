package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/galazat/cb_currency/internal/app/commands"
)

func main() {
	var source_data string

	in := bufio.NewReader(os.Stdin)
	source_data, err := in.ReadString('\n')
	if err != nil {
		log.Println(err)
	}

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	words := strings.Split(source_data, "=")
	code := strings.Split(words[1], " ")[0]
	data := words[2]

	currencies := commands.GetCurrencies(data)
	fmt.Fprint(out, currencies)

	res := currencies.Get(code)

	outputMsgText := fmt.Sprintf("%s (%s): %s", res.CharCode, res.Name, res.Value)

	fmt.Fprint(out, outputMsgText)
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
