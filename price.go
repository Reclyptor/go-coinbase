package coinbase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Price struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

// https://developers.coinbase.com/api/v2#prices
type priceService struct {
	client *CoinbaseClient
}

// https://developers.coinbase.com/api/v2#get-buy-price
func (price priceService) GetBuyPrice(currencyPair CurrencyPair) Price {
	endpoint := price.client.baseURL + "/prices" + fmt.Sprintf("/%s", currencyPair) + "/buy"
	res, _ := price.client.client.Get(endpoint)
	data, _ := ioutil.ReadAll(res.Body)
	payload := struct {
		Data Price `json:"data"`
	}{}
	_ = json.Unmarshal(data, &payload)
	return payload.Data
}

// https://developers.coinbase.com/api/v2#get-sell-price
func (price priceService) GetSellPrice(currencyPair CurrencyPair) Price {
	endpoint := price.client.baseURL + "/prices" + fmt.Sprintf("/%s", currencyPair) + "/sell"
	res, _ := price.client.client.Get(endpoint)
	data, _ := ioutil.ReadAll(res.Body)
	payload := struct {
		Data Price `json:"data"`
	}{}
	_ = json.Unmarshal(data, &payload)
	return payload.Data
}

// https://developers.coinbase.com/api/v2#get-spot-price
func (price priceService) GetSpotPrice(currencyPair CurrencyPair) Price {
	endpoint := price.client.baseURL + "/prices" + fmt.Sprintf("/%s", currencyPair) + "/spot"
	res, _ := price.client.client.Get(endpoint)
	data, _ := ioutil.ReadAll(res.Body)
	payload := struct {
		Data Price `json:"data"`
	}{}
	_ = json.Unmarshal(data, &payload)
	return payload.Data
}

// https://developers.coinbase.com/api/v2#get-spot-price
func (price priceService) GetSpotPriceHistorical(currencyPair CurrencyPair, year int, month int, day int) Price {
	endpoint := price.client.baseURL + "/prices" + fmt.Sprintf("/%s", currencyPair) + "/spot" + fmt.Sprintf("?date=%04d-%02d-%02d", year, month, day)
	res, _ := price.client.client.Get(endpoint)
	data, _ := ioutil.ReadAll(res.Body)
	payload := struct {
		Data Price `json:"data"`
	}{}
	_ = json.Unmarshal(data, &payload)
	return payload.Data
}
