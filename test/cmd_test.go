/*
 * Copyright (C) 2018 Ali Shanaakh, ashanaakh@gmail.com
 * This software may be modified and distributed under the terms
 * of the MIT license. See the LICENSE file for details.
 */

package main

import (
	"github.com/ashanaakh/crypto-rates/cmd"
	"sync"
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
	result := make([]cmd.CryptoCurrency, 0)

	var wg sync.WaitGroup
	var mutex sync.Mutex

	coins := []string{"ION", "BCH"}

	wg.Add(2)

	for _, coin := range coins {
		go cmd.GetCoinRate(coin, "EUR", &result, &wg, &mutex)
	}

	wg.Wait()

	if len(result) != 2 {
		t.Error("Expected 2 objects, got ", len(result))
	}
}
