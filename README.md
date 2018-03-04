[travis]: https://travis-ci.org/ashanaakh/crypto-rates
[goreport]: https://goreportcard.com/report/github.com/ashanaakh/crypto-rates
[license]: ./LICENSE

# Cryptocurrencies rates
[![Build Status](https://travis-ci.org/ashanaakh/crypto-rates.svg?branch=master)][travis]
[![Go Report Card](https://goreportcard.com/badge/github.com/ashanaakh/cryptio-api)][goreport]

> Binary for getting coins rates

## Overview

Thanks [chasing coins](https://chasing-coins.com) for providing awesome API for crypto-currencies rates.

## Installation

**Prerequisites**

1. Glide;
2. `$GOPATH/bin` directory in the `PATH` environment variable;

Downloads source code

```
$ go get github.com/ashanaakh/crypto-rates
```

Install dependencies

```
$ glide up
```

Compiles executable and move it to `$GOPATH/bin` directory

```
$ go install
```

## Usage

Execute binary

```
$ crypto-rates
```

*Example output:*

```
BTC 8387.85
ETH 842.728
XRP 1.06239
```
Lines can be red or green.
When line is green it means that at last hour
coin rate increased, otherwice coin rate decreased.


You have ability to choose fiat, if you want to specify fiat just use this syntax

```
$ crypto-rates -fiat=INR
```

## Development

Contributions are welcome

TODO:
- [x] Support for `BCH`,`LTC`, `DASH`, `NEO`, `ETC`, `IOT`, `EOS`
- [x] Ability to specify currencies
- [x] Ability to choose fiat
- [x] Add concurrency
- [ ] Add some wiki

## License
Project released under the terms of the MIT [license][license].
