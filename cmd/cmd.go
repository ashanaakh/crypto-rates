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
  "errors"
  "github.com/fatih/color"
  "strconv"
  "sync"
  "sort"
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
  Name   string
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
// of convertation specified coin into fiat
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
func PrettyShow(result []CryptoCurrency, fiat string) {
  color.Cyan("Fiat: " + fiat)

  sort.Slice(result, func(i, j int) bool {
    first, err := strconv.ParseFloat(result[i].Price, 64)
    second, err := strconv.ParseFloat(result[j].Price, 64)

    handleError(err)

    return first > second
  })

  for _, coin := range result {
    hour, err := strconv.ParseFloat(coin.Price, 64)

    handleError(err)

    fmt.Print(coin.Name + "\t")

    if hour > 0 {
      color.Green(coin.Price)
    } else {
      color.Red(coin.Price)
    }
  }
}

// GetCoinRate returns rate of specified coin
func GetCoinRate(code, fiat string, result *[]CryptoCurrency, wg *sync.WaitGroup, mutex *sync.Mutex) {
  url, _ := GetCoinURL(code)

  res, err := http.Get(url)

  handleError(err)

  body, err := ioutil.ReadAll(res.Body)

  handleError(err)

  coin := new(CryptoCurrency)
  coin.Name = code

  json.Unmarshal([]byte(body), &coin)

  if fiat != defaultFiat {
    convertURL, err := GetCoinConvertURL(code, fiat)

    handleError(err)

    convertRes, err := http.Get(convertURL)

    handleError(err)

    convertBody, err := ioutil.ReadAll(convertRes.Body)

    handleError(err)

    coinConvert := new(CoinConvert)
    json.Unmarshal([]byte(convertBody), &coinConvert)

    coin.Price = fmt.Sprint(coinConvert.Result)
  }

  mutex.Lock()
  *result = append(*result, *coin)
  mutex.Unlock()

  wg.Done()
}

type logic struct {
  wg sync.WaitGroup
  mutex sync.Mutex
}

// Run executes algorithm
func Run() {
  fiat := flag.String("fiat", defaultFiat, "Fiat currency")
  flag.Parse()

  result := make([]CryptoCurrency, 0)

  var wg sync.WaitGroup
  var mutex sync.Mutex

  wg.Add(len(coins))

  for _, coin := range coins {
    go GetCoinRate(coin, *fiat, &result, &wg, &mutex)
  }

  wg.Wait()

  PrettyShow(result, *fiat)
}
