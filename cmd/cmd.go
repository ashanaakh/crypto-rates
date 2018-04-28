/*
* Copyright (C) 2018 Ali Shanaakh, ashanaakh@gmail.com
* This software may be modified and distributed under the terms
* of the MIT license. See the LICENSE file for details.
 */

// Package cmd implements methods for getting latest rates for crypto-currencies 9(coins)
package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"errors"
	"github.com/fatih/color"
	"strconv"
)

// Contains checks if item exist in array
func Contains(array []string, x string) bool {
	for _, elem := range array {
		if elem == x {
			return true
		}
	}
	return false
}

// handleError handles error
func handleError(err error) {
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
}

// CoinConvert pareses json response from third party REST JSON API
type CoinConvert struct {
	Currency string      `json:"currency"`
	Result   interface{} `json:"result"`
	Coin     string      `json:"coin"`
}

// CoinChange parses json response from third party REST JSON API
type CoinChange struct {
	Hour string `json:"hour"`
	Day  string `json:"day"`
}

// CryptoCurrency for third party REST JSON API
type CryptoCurrency struct {
	Price  string     `json:"price"`
	Change CoinChange `json:"change"`
}

// Default fiat of third party REST JSON API
var defaultFiat = "USD"

// Description: List of support fiats.
var fiats = []string{"USD", "EUR", "GBP", "AUD", "CAD",
	"CNY", "EGP", "HKD", "INR", "ILS",
	"JPY", "MXP", "NZD", "PKR", "PHP",
	"RUR", "SGD", "ZAR", "KRW", "THB"}

// Description: List of support coins.
// Notes: Add any coin from list: https://chasing-coins.com/coins.
var coins = []string{"BTC", "ETH", "XRP", "DASH", "ION",
	"BCH", "LTC", "NEO", "ETC", "EOS"}

// GetCoinURL returns REST JSON API URL of specified currency rate
func GetCoinURL(coin string) (string, error) {
	if Contains(coins, coin) {
		return "https://chasing-coins.com/api/v1/std/coin/" + coin, nil
	}

	return "", errors.New("invalid coin")
}

// GetCoinConvertURL returns REST JSON API URL
// of convertion specified coin into fiat
func GetCoinConvertURL(coin, fiat string) (string, error) {
	coinOk := Contains(coins, coin)
	fiatOk := Contains(fiats, fiat)

	if coinOk && fiatOk {
		return "https://chasing-coins.com/api/v1/convert/" + coin + "/" + fiat, nil
	} else if fiatOk {
		return "", errors.New("invalid coin")
	}

	return "", errors.New("invalid fiat")
}

// PrettyShow prints rates coins color, depending on their rates
func PrettyShow(output, fiat string) {
	lines := strings.Split(output, "\n")

	color.Cyan("Fiat: " + fiat)

	for _, line := range lines {
		parts := strings.Split(line, " ")

		if len(parts) < 3 {
			continue
		}

		coin := parts[0]
		price := parts[1]
		hour, _ := strconv.ParseFloat(parts[2], 64)

		if hour > 0 {
			color.Green(coin + " " + price)
		} else {
			color.Red(coin + " " + price)
		}
	}
}

// GetCoinRate returns rate of specified coin
func GetCoinRate(coinName, fiat string, msg chan string) {
	url, _ := GetCoinURL(coinName)

	res, err := http.Get(url)

	handleError(err)

	body, err := ioutil.ReadAll(res.Body)

	handleError(err)

	coin := new(CryptoCurrency)
	json.Unmarshal([]byte(body), &coin)

	if fiat != defaultFiat {
		convertURL, _ := GetCoinConvertURL(coinName, fiat)
		convertRes, err := http.Get(convertURL)

		handleError(err)

		convertBody, err := ioutil.ReadAll(convertRes.Body)

		handleError(err)

		coinConvert := new(CoinConvert)
		json.Unmarshal([]byte(convertBody), &coinConvert)

		coin.Price = fmt.Sprint(coinConvert.Result)
	}

	msg <- coinName + " " + coin.Price + " " + coin.Change.Hour + "\n"
}

// Run executes algorithm
func Run() {
	fiat := flag.String("fiat", defaultFiat, "Fiat currency")
	flag.Parse()

	var output string

	ch := make(chan string, len(coins))

	for _, coin := range coins {
		go GetCoinRate(coin, *fiat, ch)
	}

	for range coins {
		output += <-ch
	}

	PrettyShow(output, *fiat)
}
