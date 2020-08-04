package coinbase

import (
	"encoding/json"
	"io/ioutil"
)

type Currency struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	MinimumSize string `json:"min_size"`

	// WebSocket Only
	Status        string `json:"status"`
	StatusMessage string `json:"status_message"`

}

// https://developers.coinbase.com/api/v2#currencies
type currencyService struct {
	client *CoinbaseClient
}

// https://docs.pro.coinbase.com/#currencies
type currencyProService struct {
	client *CoinbaseProClient
}

// https://developers.coinbase.com/api/v2#get-currencies
func (currency currencyService) GetCurrencies() []Currency {
	endpoint := currency.client.baseURL + "/currencies"
	res, _ := currency.client.client.Get(endpoint)
	data, _ := ioutil.ReadAll(res.Body)
	payload := struct {
		Data []Currency `json:"data"`
	}{}
	_ = json.Unmarshal(data, &payload)
	return payload.Data
}

// https://docs.pro.coinbase.com/#get-currencies
func (currency currencyProService) GetCurrencies() []Currency {
	endpoint := currency.client.baseURL + "/currencies"
	res, _ := currency.client.client.Get(endpoint)
	data, _ := ioutil.ReadAll(res.Body)
	var payload []Currency
	_ = json.Unmarshal(data, &payload)
	return payload
}
