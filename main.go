package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/fatih/color"
)

// List of supported currencies
var coins = []string{"BTC", "ETH", "XRP"}

// CryptoCurrency for api resp
type CryptoCurrency struct {
	Price string `json:"price"`
}

func getCoinURL(base string) string {
	return "https://chasing-coins.com/api/v1/std/coin/" + base
}

func printCurrenciesRates() {
	for _, coin := range coins {

		url := getCoinURL(coin)

		res, err := http.Get(url)

		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}

		body, err := ioutil.ReadAll(res.Body)

		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}

		crypt := new(CryptoCurrency)
		json.Unmarshal([]byte(body), &crypt)

		boldYellow := color.New(color.FgYellow).Add(color.Bold)
		boldYellow.Print(coin + " ")
		color.Green(crypt.Price)
	}
}

func main() {
	printCurrenciesRates()
}
