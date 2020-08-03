package coinbase

import (
	"net/http"
)

type CoinbaseClient struct {
	baseURL string

	client http.Client

	Currency currencyService
	Exchange exchangeService
	Price    priceService
	Time     timeService
}

type CoinbaseProClient struct {
	baseURL      string
	sandboxURL   string
	websocketURL string
	origin       string

	client http.Client

	Product  productService
	Currency currencyProService
	Time     timeProService
	Channel  channelService
}

func NewCoinbaseClient() CoinbaseClient {
	client := CoinbaseClient {
		baseURL: "https://api.coinbase.com/v2",
		client:  *http.DefaultClient,
	}

	client.Currency = currencyService{client: &client}
	client.Exchange = exchangeService{client: &client}
	client.Price    = priceService{client: &client}
	client.Time     = timeService{client: &client}

	return client
}

func NewCoinbaseProClient() CoinbaseProClient {
	client := CoinbaseProClient {
		baseURL:      "https://api.pro.coinbase.com",
		sandboxURL:   "https://api-public.sandbox.pro.coinbase.com",
		websocketURL: "wss://ws-feed.pro.coinbase.com",
		origin:       "https://coinbase.com",
		client:       *http.DefaultClient,
	}

	client.Product  = productService{client: &client}
	client.Currency = currencyProService{client: &client}
	client.Time     = timeProService{client: &client}
	client.Channel  = channelService{client: &client}

	return client
}
