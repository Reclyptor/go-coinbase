package coinbase

import (
	"encoding/json"
	"io/ioutil"
)

type Time struct {
	ISO   string  `json:"iso"`
	Epoch float64 `json:"epoch"`
}

// https://developers.coinbase.com/api/v2#time
type timeService struct {
	client *CoinbaseClient
}

// https://docs.pro.coinbase.com/#time
type timeProService struct {
	client *CoinbaseProClient
}

// https://developers.coinbase.com/api/v2#get-current-time
func (time *timeService) GetCurrentTime() Time {
	endpoint := time.client.baseURL + "/time"
	res, _ := time.client.client.Get(endpoint)
	data, _ := ioutil.ReadAll(res.Body)
	payload := struct {
		Data Time `json:"data"`
	}{}
	_ = json.Unmarshal(data, &payload)
	return payload.Data
}

// https://docs.pro.coinbase.com/#time
func (time *timeProService) GetCurrentTime() Time {
	endpoint := time.client.baseURL + "/time"
	res, _ := time.client.client.Get(endpoint)
	data, _ := ioutil.ReadAll(res.Body)
	var payload Time
	_ = json.Unmarshal(data, &payload)
	return payload
}
