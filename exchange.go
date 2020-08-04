package coinbase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ExchangeRates struct {
	Currency string            `json:"currency"`
	Rates    map[string]string `json:"rates"`
}

// https://developers.coinbase.com/api/v2#exchange-rates
type exchangeService struct {
	client *CoinbaseClient
}

// https://developers.coinbase.com/api/v2#get-exchange-rates
func (exchange *exchangeService) GetExchangeRates(crypto CryptoCurrency) ExchangeRates {
	endpoint := exchange.client.baseURL + "/exchange-rates" + fmt.Sprintf("?currency=%s", crypto)
	res, _ := exchange.client.client.Get(endpoint)
	data, _ := ioutil.ReadAll(res.Body)
	payload := struct {
		Data ExchangeRates `json:"data"`
	}{}
	_ = json.Unmarshal(data, &payload)
	return payload.Data
}
