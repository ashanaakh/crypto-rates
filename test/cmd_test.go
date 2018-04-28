package main

import (
  "github.com/ashanaakh/crypto-rates/cmd"
  "strings"
  "testing"
)

func TestGetValidCoinURL(t *testing.T) {
  url, _ := cmd.GetCoinURL("BTC")
  expectedURL := "https://chasing-coins.com/api/v1/std/coin/BTC"

  if url != expectedURL {
    t.Error("Expected"+expectedURL+", got ", url)
  }
}

func TestGetInvalidCoinURL(t *testing.T) {
  url, err := cmd.GetCoinURL("NOT VALID COIN")
  expectedError := "invalid coin"

  if url != "" {
    t.Error("Expected"+"empty string, got ", url)
  }

  if err == nil {
    t.Error("Expected error: "+err.Error()+", got ", err)
  }

  if err.Error() != expectedError {
    t.Error("Expected error: "+err.Error()+", got ", err.Error())
  }
}

func TestGetValidGetCoinConvertURL(t *testing.T) {
  url, _ := cmd.GetCoinConvertURL("ION", "INR")
  expectedURL := "https://chasing-coins.com/api/v1/convert/ION/INR"

  if url != expectedURL {
    t.Error("Expected"+expectedURL+", got ", url)
  }
}

func TestGetInvalidGetCoinConvertURL(t *testing.T) {
  url, err := cmd.GetCoinConvertURL("NOT VALID COIN", "INR")
  expectedCoinError := "invalid coin"

  if url != "" {
    t.Error("Expected"+"empty string, got ", url)
  }

  if err == nil {
    t.Error("Expected error: "+err.Error()+", got ", err)
  }

  if err.Error() != expectedCoinError {
    t.Error("Expected error: "+err.Error()+", got ", err.Error())
  }

  url, err = cmd.GetCoinConvertURL("ETH", "NOT VALID COIN")
  expectedFiatError := "invalid fiat"

  if url != "" {
    t.Error("Expected"+"empty string, got ", url)
  }

  if err == nil {
    t.Error("Expected error: "+err.Error()+", got ", err)
  }

  if err.Error() != expectedFiatError {
    t.Error("Expected error: "+err.Error()+", got ", err.Error())
  }
}

func TestValidContains(t *testing.T) {
  array := []string{"VALID"}
  validItem := "VALID"
  invalidItem := "INVALID"

  if !cmd.Contains(array, validItem) {
    t.Error("Expected true, got false")
  }

  if cmd.Contains(array, invalidItem) {
    t.Error("Expected false, got true")
  }
}

func TestGetCoinRate(t *testing.T) {
  var output string

  coins := []string{"ION", "BCH"}

  ch := make(chan string, 2)

  for _, coin := range coins {
    go cmd.GetCoinRate(coin, "INR", ch)
  }

  for range coins {
    output += <-ch
  }

  lines := strings.Split(output, "\n")

  if len(lines) != 3 {
    t.Error("Expected 3 lines, got ", len(lines))
  }
}
