/*
 * Copyright (C) 2018 Ali Shanaakh, ashanaakh@gmail.com
 //* This software may be modified and distributed under the terms
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
  "strconv"
  "github.com/fatih/color"
)

// cryptoCurrencyConvert pareses request json, when fiat isn't USD
type cryptoCurrencyConvert struct {
  Currency string `json:"currency"`
  Result interface{}   `json:"result"`
  Coin string     `json:"coin"`
}

// cryptoCurrencyChange
type cryptoCurrencyChange struct {
	Hour string `json:"hour"`
	Day  string `json:"day"`
}

// cryptoCurrency for api resp
type cryptoCurrency struct {
	Price  string               `json:"price"`
	Change cryptoCurrencyChange `json:"change"`
}

// fiats [Array<String>]
//
// Description: List of support fiats.
var fiats = []string{ "USD", "EUR", "GBP", "AUD", "CAD",
                      "CNY", "EGP", "HKD", "INR", "ILS",
                      "JPY", "MXP", "NZD", "PKR", "PHP",
                      "RUR", "SGD", "ZAR", "KRW", "THB" }

// coins [Array<String>]
//
// Description: List of support coins.
//
// Notes: Add any coin from list: https://chasing-coins.com/coins.
var coins = []string{ "BTC", "ETH", "XRP", "DASH", "ION",
                      "BCH", "LTC", "NEO", "ETC", "EOS" }

// getCoinURL returns URL of REST JSON API of specified currency rate
func getCoinURL(coin string) string {
	return "https://chasing-coins.com/api/v1/std/coin/" + coin
}

func getCoinConvertURL(coin, fiat string) string {
  return "https://chasing-coins.com/api/v1/convert/" + coin + "/" + fiat
}

// handleError handles error
func handleError(err error) {
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
}

// showCoinRate prints rate of specified coin
func showCoinRate(coin, fiat string, msg chan string) {
  url := getCoinURL(coin)

  res, err := http.Get(url)

  handleError(err)

  body, err := ioutil.ReadAll(res.Body)

  handleError(err)

  crypt := new(cryptoCurrency)
  json.Unmarshal([]byte(body), &crypt)

  hour, _ := strconv.ParseFloat(crypt.Change.Hour, 64)

  if fiat != "USD" {
    convertURL := getCoinConvertURL(coin, fiat)
    convertRes, err := http.Get(convertURL)

    handleError(err)

    convertBody, err := ioutil.ReadAll(convertRes.Body)

    handleError(err)

    coinConvert := new(cryptoCurrencyConvert)
    json.Unmarshal([]byte(convertBody), &coinConvert)

    crypt.Price = fmt.Sprint(coinConvert.Result)
  }

  if hour > 0 {
    color.Green(coin + " " + crypt.Price)
  } else {
    color.Red(coin + " " + crypt.Price)
  }

  msg <- ""
}

// showCoinRate prints rates of all cryoto-currencies in coins array
func showCoinsRates(fiat string) {

  result := make(chan string, len(coins))

	for _, coin := range coins {
		go showCoinRate(coin, fiat, result)
	}

	for range coins {
    _ = <-result
  }
}

// Run executes algorithm
func Run() {
	fiat := flag.String("fiat", "USD", "Fiat currency")
  flag.Parse()

  fmt.Println("fiat: ", *fiat)

  showCoinsRates(*fiat)
}
