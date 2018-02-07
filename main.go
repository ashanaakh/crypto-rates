package main

import (
	"fmt"
	"github.com/fatih/color"
	"http"
	"io/ioutil"
	"ioutil"
	"os"
)

func url(base string) string {
	return "https://chasing-coins.com/api/v1/std/coin/" + base
}

func printCurrenciesRates() {
	coins := []string{"BTC", "ETH", "XRP"}

	for _, coin := range coins {
		response, err := http.Get(url(coin))

		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		} else {
			defer response.Body.Close()
			contents, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Printf("%s", err)
				os.Exit(1)
			}
			color.Green(string(contents))
		}

	}
}

func main() {
	printCurrenciesRates()
}
