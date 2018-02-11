/*
 * Copyright (C) 2018 Ali Shanaakh, ashanaakh@gmail.com
 * This software may be modified and distributed under the terms
 * of the MIT license. See the LICENSE file for details.
 */

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/fatih/color"
)

// List of supported currencies
var coins = []string{"BTC", "ETH", "XRP"}

// cryptoCurrencyChange
type cryptoCurrencyChange struct {
	Hour string `json:"hour"`
	Day  string `json:"day"`
}

// CryptoCurrency for api resp
type cryptoCurrency struct {
	Price  string               `json:"price"`
	Change cryptoCurrencyChange `json:"change"`
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

		crypt := new(cryptoCurrency)
		json.Unmarshal([]byte(body), &crypt)

		boldYellow := color.New(color.FgYellow).Add(color.Bold)
		boldYellow.Print(coin + " ")

		hour, _ := strconv.ParseFloat(crypt.Change.Hour, 64)

		if hour > 0 {
			color.Green(crypt.Price)
		} else {
			color.Red(crypt.Price)
		}

	}
}

func main() {
	printCurrenciesRates()
}
