[travis]: https://travis-ci.org/ashanaakh/crypto-rates
[goreport]: https://goreportcard.com/report/github.com/ashanaakh/crypto-rates
[license]: ./LICENSE

# Crypto rates
[![Build Status](https://travis-ci.org/ashanaakh/crypto-rates.svg?branch=master)][travis]
[![Go Report Card](https://goreportcard.com/badge/github.com/ashanaakh/cryptio-api)][goreport]

> Binary for getting coins rates

## Overview

Thanks [chasing coins](https://chasing-coins.com) for providing awesome API for crypto-currencies rates.

## Installation

**Prerequisites**

- Glide.
- `$GOPATH/bin` directory in the `PATH` environment variable.

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
Fiat: USD
BTC 9459.94
BCH 1444.9
ETH 693.658
DASH 499.435
LTC 154.333
NEO 80.3039
EOS 22.8904
ETC 21.7837
ION 2.91656
XRP 0.885313
```
Lines can be red or green.
When line is green it means that at last hour
coin rate increased, otherwice coin rate decreased.


You have ability to choose fiat, if you want to specify fiat just use this syntax.

```
$ crypto-rates -fiat=INR
```

## Development

Contributions are welcome.

## License
Project released under the terms of the MIT [license][license].
