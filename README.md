[travis]: https://travis-ci.org/shal/crypto-rates
[goreport]: https://goreportcard.com/report/github.com/shal/crypto-rates
[license]: ./LICENSE

# Crypto rates
[![Build Status](https://travis-ci.org/shal/crypto-rates.svg?branch=master)][travis]
[![Go Report Card](https://goreportcard.com/badge/github.com/shal/cryptio-api)][goreport]

> Tool for getting coins rates

## Overview
## Installation
Install using go dep

```
$ go get github.com/shal/crypto-rates
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
Download source code

```
$ go get github.com/shal/crypto-rates
```

Install dependencies

```
$ go get ./...
```

Compile executable and add it to `$GOPATH/bin` directory

```
$ go install
```

Contributions are welcome.

Thanks [chasing coins](https://chasing-coins.com) for providing awesome API for crypto-currencies rates.

## License
Project released under the terms of the MIT [license][license].
