# Cryptocurrencies rates
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
- Support for `BCH`,`LTC`, `DASH`, `NEO`, `ETC`, `IOT`, `EOS`
- Opportunity to specify currencies
- Opportunity to choose fiat

## License
Project released under the terms of the MIT [license](LICENSE).
