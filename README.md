# go-coinbase
Coinbase API Client \
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/Reclyptor/go-coinbase?color=blue&label=Release&sort=semver&style=plastic)
![GitHub](https://img.shields.io/github/license/Reclyptor/go-coinbase?color=red&label=License&style=plastic)
![GitHub repo size](https://img.shields.io/github/repo-size/Reclyptor/go-coinbase?color=green&label=Size&style=plastic)

## Installation
```shell script
go get -d -u github.com/Reclyptor/go-coinbase
```

## Sample Usage
```go
package main

import (
	"fmt"
	"github.com/Reclyptor/go-coinbase"
)

func main() {
	// Create An API Client
	client := coinbase.NewCoinbaseClient()

	// Retrieve An Asset's Market Price
	price := client.Price.GetSpotPrice(coinbase.CurrencyPairs.BTC_USD)

	fmt.Printf("%+v\n", price)
}
```

## API
The official API documentation can be found on the Coinbase developer portal. ([Base](https://developers.coinbase.com/api/v2]) | [Pro](https://docs.pro.coinbase.com/#api))

The endpoints supported by the client are listed below.

### Base
- [Data Endpoints](https://developers.coinbase.com/api/v2#data-endpoints)
  - [Currencies](https://developers.coinbase.com/api/v2#currencies)
    - [Get Currencies](https://developers.coinbase.com/api/v2#get-currencies)
  - [Exchange Rates](https://developers.coinbase.com/api/v2#exchange-rates)
    - [Get Exchange Rates](https://developers.coinbase.com/api/v2#get-exchange-rates)
  - [Prices](https://developers.coinbase.com/api/v2#prices)
    - [Get Buy Price](https://developers.coinbase.com/api/v2#get-buy-price)
    - [Get Sell Price](https://developers.coinbase.com/api/v2#get-sell-price)
    - [Get Spot Price](https://developers.coinbase.com/api/v2#get-spot-price)
  - [Time](https://developers.coinbase.com/api/v2#time)
    - [Get Current Time](https://developers.coinbase.com/api/v2#get-current-time)

### Pro
- [Market Data](https://docs.pro.coinbase.com/#market-data)
  - [Products](https://docs.pro.coinbase.com/#products)
    - [Get Products](https://docs.pro.coinbase.com/#get-products)
    - [Get Single Product](https://docs.pro.coinbase.com/#get-single-product)
    - [Get Product Order Book](https://docs.pro.coinbase.com/#get-product-order-book)
    - [Get Product Ticker](https://docs.pro.coinbase.com/#get-product-ticker)
    - [Get Trades](https://docs.pro.coinbase.com/#get-trades)
    - [Get Historic Rates](https://docs.pro.coinbase.com/#get-historic-rates)
    - [Get 24hr Stats](https://docs.pro.coinbase.com/#get-24hr-stats)
  - [Currencies](https://docs.pro.coinbase.com/#currencies)
    - [Get Currencies](https://docs.pro.coinbase.com/#get-currencies)
  - [Time](https://docs.pro.coinbase.com/#time)
    - [Get Current Time](https://docs.pro.coinbase.com/#time)
   
## Roadmap
1) Formal Testing
2) Godoc
3) Proper Error Handling and Propagation
5) Websocket Feed
4) Support for Private Endpoints

## Notes
2020 August 03
- Laid out the initial groundwork for WebSocket support. Currently, only the Heartbeat channel is subscribable.

2020 August 02
 - This initial version of the client only supports the use of the public data endpoints, which do not require any form of authentication.
   - Note that these endpoints are throttled by Coinbase at a rate of roughly 2 requests per second.
