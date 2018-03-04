[travis]: https://travis-ci.org/ashanaakh/crypto-rates
[goreport]: https://goreportcard.com/report/github.com/ashanaakh/crypto-rates
[license]: ./LICENSE

# Cryptocurrencies rates
[![Build Status](https://travis-ci.org/ashanaakh/crypto-rates.svg?branch=master)][travis]
[![Go Report Card](https://goreportcard.com/badge/github.com/ashanaakh/cryptio-api)][goreport]


> Command line interface for getting latest crypto rates

## Overview

Thanks [chasing coins](https://chasing-coins.com) for providing awesome API for latest cryptocurrencies rates.

## Installation

**Prerequisites**

1. Glide
2. `$GOPATH/bin` directory in the `PATH` environment variable

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

## Development

Contributions are welcome

TODO:
- [ ] Support for `BCH`,`LTC`, `DASH`, `NEO`, `ETC`, `IOT`, `EOS`
- [ ] Ability to specify currencies
- [ ] Ability to choose fiat

## License
Project released under the terms of the MIT [license][license].
