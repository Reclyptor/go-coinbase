package coinbase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"time"
)

type Product struct {
	ID              CurrencyPair `json:"id"`
	DisplayName     string       `json:"display_name"`
	BaseCurrency    string       `json:"base_currency"`
	QuoteCurrency   string       `json:"quote_currency"`
	BaseIncrement   string       `json:"base_increment"`
	QuoteIncrement  string       `json:"quote_increment"`
	BaseMinSize     string       `json:"base_min_size"`
	BaseMaxSize     string       `json:"base_max_size"`
	MinMarketFunds  string       `json:"min_market_funds"`
	MaxMarketFunds  string       `json:"max_market_funds"`
	Status          string       `json:"status"`
	StatusMessage   string        `json:"status_message"`
	CancelOnly      bool          `json:"cancel_only"`
	LimitOnly       bool          `json:"limit_only"`
	PostOnly        bool          `json:"post_only"`
	TradingDisabled bool          `json:"trading_disabled"`
}

type Book struct {
	Sequence int64 `json:"sequence"`
	Bids     []Bid `json:"bids"`
	Asks     []Ask `json:"asks"`
}

type Ticker struct {
	TradeID int64  `json:"trade_id"`
	Price   string `json:"price"`
	Size    string `json:"size"`
	Bid     string `json:"bid"`
	Ask     string `json:"ask"`
	Volume  string `json:"volume"`
	Time    string `json:"time"`
}

type Trade struct {
	Time    string `json:"time"`
	TradeID int64  `json:"trade_id"`
	Price   string `json:"price"`
	Size    string `json:"size"`
	Side    string `json:"side"`
}

type Stats struct {
	Open        string `json:"open"`
	High        string `json:"high"`
	Low         string `json:"low"`
	Volume      string `json:"volume"`
	Last        string `json:"last"`
	Volume30Day string `json:"volume_30day"`
}

type Bid []interface{}

type Ask []interface{}

type Candle []interface{}

// https://docs.pro.coinbase.com/#products
type productService struct {
	client *CoinbaseProClient
}

// https://docs.pro.coinbase.com/#get-products
func (product productService) GetProducts() []Product {
	endpoint := product.client.baseURL + "/products"
	res, _ := product.client.client.Get(endpoint)
	data, _ := ioutil.ReadAll(res.Body)
	var payload []Product
	_ = json.Unmarshal(data, &payload)
	return payload
}

// https://docs.pro.coinbase.com/#get-single-product
func (product productService) GetSingleProduct(currencyPair CurrencyPair) Product {
	endpoint := product.client.baseURL + "/products" + fmt.Sprintf("/%s", currencyPair)
	res, _ := product.client.client.Get(endpoint)
	data, _ := ioutil.ReadAll(res.Body)
	var payload Product
	_ = json.Unmarshal(data, &payload)
	return payload
}


// https://docs.pro.coinbase.com/#get-product-order-book
func (product productService) GetProductOrderBook(currencyPair CurrencyPair, level int) Book {
	endpoint := product.client.baseURL + "/products" + fmt.Sprintf("/%s", currencyPair) + "/book" + fmt.Sprintf("?level=%d", int(math.Max(math.Min(float64(level), 3), 1)))
	res, _ := product.client.client.Get(endpoint)
	data, _ := ioutil.ReadAll(res.Body)
	var payload Book
	_ = json.Unmarshal(data, &payload)
	return payload
}

// https://docs.pro.coinbase.com/#get-product-ticker
func (product productService) GetProductTicker(currencyPair CurrencyPair) Ticker {
	endpoint := product.client.baseURL + "/products" + fmt.Sprintf("/%s", currencyPair) + "/ticker"
	res, _ := product.client.client.Get(endpoint)
	data, _ := ioutil.ReadAll(res.Body)
	var payload Ticker
	_ = json.Unmarshal(data, &payload)
	return payload
}

// https://docs.pro.coinbase.com/#get-trades
func (product productService) GetTrades(currencyPair CurrencyPair) []Trade {
	endpoint := product.client.baseURL + "/products" + fmt.Sprintf("/%s", currencyPair) + "/trades"
	res, _ := product.client.client.Get(endpoint)
	data, _ := ioutil.ReadAll(res.Body)
	var payload []Trade
	_ = json.Unmarshal(data, &payload)
	return payload
}

// https://docs.pro.coinbase.com/#get-historic-rates
func (product productService) GetHistoricRates(currencyPair CurrencyPair, start time.Time, end time.Time, granularity time.Duration) []Candle {
	endpoint := product.client.baseURL + "/products" + fmt.Sprintf("/%s", currencyPair) + "/candles" + fmt.Sprintf("?start=%s&end=%s&granularity=%s", start.Format(time.RFC3339), end.Format(time.RFC3339), granularity.Seconds())
	res, _ := product.client.client.Get(endpoint)
	data, _ := ioutil.ReadAll(res.Body)
	var payload []Candle
	_ = json.Unmarshal(data, &payload)
	return payload
}

// https://docs.pro.coinbase.com/#get-24hr-stats
func (product productService) Get24HourStats(currencyPair CurrencyPair) Stats {
	endpoint := product.client.baseURL + "/products" + fmt.Sprintf("/%s", currencyPair) + "/stats"
	res, _ := product.client.client.Get(endpoint)
	data, _ := ioutil.ReadAll(res.Body)
	var payload Stats
	_ = json.Unmarshal(data, &payload)
	return payload
}

func (bid Bid) Price() float64 {
	if len(bid) < 1 {
		return 0
	}

	price, err := strconv.ParseFloat(bid[0].(string), 64)

	if err != nil {
		return 0
	}

	return price
}

func (bid Bid) Size() float64 {
	if len(bid) < 2 {
		return 0
	}

	size, err := strconv.ParseFloat(bid[1].(string), 64)

	if err != nil {
		return 0
	}

	return size
}

func (bid Bid) NumOrders() int64 {
	if len(bid) < 3 {
		return 0
	}

	if numOrders, ok := bid[2].(float64); ok {
		return int64(numOrders)
	}

	return 0
}

func (bid Bid) OrderID() string {
	if len(bid) < 3 {
		return ""
	}

	if orderID, ok := bid[2].(string); ok {
		return orderID
	}

	return ""
}

func (ask Ask) Price() float64 {
	if len(ask) < 1 {
		return 0
	}

	price, err := strconv.ParseFloat(ask[0].(string), 64)

	if err != nil {
		return 0
	}

	return price
}

func (ask Ask) Size() float64 {
	if len(ask) < 2 {
		return 0
	}

	size, err := strconv.ParseFloat(ask[1].(string), 64)

	if err != nil {
		return 0
	}

	return size
}

func (ask Ask) NumOrders() int64 {
	if len(ask) < 3 {
		return 0
	}

	if numOrders, ok := ask[2].(float64); ok {
		return int64(numOrders)
	}

	return 0
}

func (ask Ask) OrderID() string {
	if len(ask) < 3 {
		return ""
	}

	if orderID, ok := ask[2].(string); ok {
		return orderID
	}

	return ""
}

func (candle Candle) Time() int64 {
	if len(candle) < 1 {
		return 0
	}

	if time, ok := candle[0].(float64); ok {
		return int64(time)
	}

	return 0
}

func (candle Candle) Low() float64 {
	if len(candle) < 2 {
		return 0
	}

	if low, ok := candle[1].(float64); ok {
		return low
	}

	return 0
}

func (candle Candle) High() float64 {
	if len(candle) < 3 {
		return 0
	}

	if high, ok := candle[2].(float64); ok {
		return high
	}

	return 0
}

func (candle Candle) Open() float64 {
	if len(candle) < 4 {
		return 0
	}

	if open, ok := candle[3].(float64); ok {
		return open
	}

	return 0
}

func (candle Candle) Close() float64 {
	if len(candle) < 5 {
		return 0
	}

	if close_, ok := candle[4].(float64); ok {
		return close_
	}

	return 0
}

func (candle Candle) Volume() float64 {
	if len(candle) < 6 {
		return 0
	}

	if volume, ok := candle[5].(float64); ok {
		return volume
	}

	return 0
}
