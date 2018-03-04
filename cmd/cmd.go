/*
 * Copyright (C) 2018 Ali Shanaakh, ashanaakh@gmail.com
 * This software may be modified and distributed under the terms
 * of the MIT license. See the LICENSE file for details.
 */

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

var coins = []string{"BTC", "ETH", "XRP", "DASH", "ION"}

func getCoinURL(currency string) string {
	return "https://chasing-coins.com/api/v1/std/coin/" + currency
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
}

// func getCoinColor(percent int) (cf func(string, ...interface{})) {
// 	if percent > 0 {
// 		cf = color.Green
// 	} else {
// 		cf = color.Red
// 	}
// 	return cf
// }

func printCurrenciesRates() {
	for _, coin := range coins {
		url := getCoinURL(coin)
		res, err := http.Get(url)

		handleError(err)

		body, err := ioutil.ReadAll(res.Body)

		handleError(err)

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

// Run - main function in cmd package, that execute algorithm
func Run() {
	fiat := flag.String("fiat", "USD", "Fiat currency")

	flag.Parse()

	fmt.Println("fiat: ", *fiat)

	printCurrenciesRates()
}
