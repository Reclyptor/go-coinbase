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

type CryptoCurrency string
const (
	ALGO CryptoCurrency = "ALGO"
	ATOM CryptoCurrency = "ATOM"
	BAT  CryptoCurrency = "BAT"
	BTC  CryptoCurrency = "BTC"
	BCH  CryptoCurrency = "BCH"
	BSV  CryptoCurrency = "BSV"
	COMP CryptoCurrency = "COMP"
	CVC  CryptoCurrency = "CVC"
	DAI  CryptoCurrency = "DAI"
	DASH CryptoCurrency = "DASH"
	DNT  CryptoCurrency = "DNT"
	EOS  CryptoCurrency = "EOS"
	ETH  CryptoCurrency = "ETH"
	ETC  CryptoCurrency = "ETC"
	GNT  CryptoCurrency = "GNT"
	KNC  CryptoCurrency = "KNC"
	LINK CryptoCurrency = "LINK"
	LOOM CryptoCurrency = "LOOM"
	LTC  CryptoCurrency = "LTC"
	MANA CryptoCurrency = "MANA"
	MKR  CryptoCurrency = "MKR"
	OMG  CryptoCurrency = "OMG"
	OXT  CryptoCurrency = "OXT"
	REP  CryptoCurrency = "REP"
	USDC CryptoCurrency = "USDC"
	XLM  CryptoCurrency = "XLM"
	XRP  CryptoCurrency = "XRP"
	XTZ  CryptoCurrency = "XTZ"
	ZEC  CryptoCurrency = "ZEC"
	ZRX  CryptoCurrency = "ZRX"
)

type FiatCurrency string
const (
	AED FiatCurrency = "AED"
	AFN FiatCurrency = "AFN"
	ALL FiatCurrency = "ALL"
	AMD FiatCurrency = "AMD"
	ANG FiatCurrency = "ANG"
	AOA FiatCurrency = "AOA"
	ARS FiatCurrency = "ARS"
	AUD FiatCurrency = "AUD"
	AWG FiatCurrency = "AWG"
	AZN FiatCurrency = "AZN"
	BAM FiatCurrency = "BAM"
	BBD FiatCurrency = "BBD"
	BDT FiatCurrency = "BDT"
	BGN FiatCurrency = "BGN"
	BHD FiatCurrency = "BHD"
	BIF FiatCurrency = "BIF"
	BMD FiatCurrency = "BMD"
	BND FiatCurrency = "BND"
	BOB FiatCurrency = "BOB"
	BRL FiatCurrency = "BRL"
	BSD FiatCurrency = "BSD"
	BTN FiatCurrency = "BTN"
	BWP FiatCurrency = "BWP"
	BYN FiatCurrency = "BYN"
	BYR FiatCurrency = "BYR"
	BZD FiatCurrency = "BZD"
	CAD FiatCurrency = "CAD"
	CDF FiatCurrency = "CDF"
	CHF FiatCurrency = "CHF"
	CLF FiatCurrency = "CLF"
	CLP FiatCurrency = "CLP"
	CNH FiatCurrency = "CNH"
	CNY FiatCurrency = "CNY"
	COP FiatCurrency = "COP"
	CRC FiatCurrency = "CRC"
	CUC FiatCurrency = "CUC"
	CVE FiatCurrency = "CVE"
	CZK FiatCurrency = "CZK"
	DJF FiatCurrency = "DJF"
	DKK FiatCurrency = "DKK"
	DOP FiatCurrency = "DOP"
	DZD FiatCurrency = "DZD"
	EEK FiatCurrency = "EEK"
	EGP FiatCurrency = "EGP"
	ERN FiatCurrency = "ERN"
	ETB FiatCurrency = "ETB"
	EUR FiatCurrency = "EUR"
	FJD FiatCurrency = "FJD"
	FKP FiatCurrency = "FKP"
	GBP FiatCurrency = "GBP"
	GEL FiatCurrency = "GEL"
	GGP FiatCurrency = "GGP"
	GHS FiatCurrency = "GHS"
	GIP FiatCurrency = "GIP"
	GMD FiatCurrency = "GMD"
	GNF FiatCurrency = "GNF"
	GTQ FiatCurrency = "GTQ"
	GYD FiatCurrency = "GYD"
	HKD FiatCurrency = "HKD"
	HNL FiatCurrency = "HNL"
	HRK FiatCurrency = "HRK"
	HTG FiatCurrency = "HTG"
	HUF FiatCurrency = "HUF"
	IDR FiatCurrency = "IDR"
	ILS FiatCurrency = "ILS"
	IMP FiatCurrency = "IMP"
	INR FiatCurrency = "INR"
	IQD FiatCurrency = "IQD"
	ISK FiatCurrency = "ISK"
	JEP FiatCurrency = "JEP"
	JMD FiatCurrency = "JMD"
	JOD FiatCurrency = "JOD"
	JPY FiatCurrency = "JPY"
	KES FiatCurrency = "KES"
	KGS FiatCurrency = "KGS"
	KHR FiatCurrency = "KHR"
	KMF FiatCurrency = "KMF"
	KRW FiatCurrency = "KRW"
	KWD FiatCurrency = "KWD"
	KYD FiatCurrency = "KYD"
	KZT FiatCurrency = "KZT"
	LAK FiatCurrency = "LAK"
	LBP FiatCurrency = "LBP"
	LKR FiatCurrency = "LKR"
	LRD FiatCurrency = "LRD"
	LSL FiatCurrency = "LSL"
	LTL FiatCurrency = "LTL"
	LVL FiatCurrency = "LVL"
	LYD FiatCurrency = "LYD"
	MAD FiatCurrency = "MAD"
	MDL FiatCurrency = "MDL"
	MGA FiatCurrency = "MGA"
	MKD FiatCurrency = "MKD"
	MMK FiatCurrency = "MMK"
	MNT FiatCurrency = "MNT"
	MOP FiatCurrency = "MOP"
	MRO FiatCurrency = "MRO"
	MTL FiatCurrency = "MTL"
	MUR FiatCurrency = "MUR"
	MVR FiatCurrency = "MVR"
	MWK FiatCurrency = "MWK"
	MXN FiatCurrency = "MXN"
	MYR FiatCurrency = "MYR"
	MZN FiatCurrency = "MZN"
	NAD FiatCurrency = "NAD"
	NGN FiatCurrency = "NGN"
	NIO FiatCurrency = "NIO"
	NOK FiatCurrency = "NOK"
	NPR FiatCurrency = "NPR"
	NZD FiatCurrency = "NZD"
	OMR FiatCurrency = "OMR"
	PAB FiatCurrency = "PAB"
	PEN FiatCurrency = "PEN"
	PGK FiatCurrency = "PGK"
	PHP FiatCurrency = "PHP"
	PKR FiatCurrency = "PKR"
	PLN FiatCurrency = "PLN"
	PYG FiatCurrency = "PYG"
	QAR FiatCurrency = "QAR"
	RON FiatCurrency = "RON"
	RSD FiatCurrency = "RSD"
	RUB FiatCurrency = "RUB"
	RWF FiatCurrency = "RWF"
	SAR FiatCurrency = "SAR"
	SBD FiatCurrency = "SBD"
	SCR FiatCurrency = "SCR"
	SEK FiatCurrency = "SEK"
	SGD FiatCurrency = "SGD"
	SHP FiatCurrency = "SHP"
	SLL FiatCurrency = "SLL"
	SOS FiatCurrency = "SOS"
	SRD FiatCurrency = "SRD"
	SSP FiatCurrency = "SSP"
	STD FiatCurrency = "STD"
	SVC FiatCurrency = "SVC"
	SZL FiatCurrency = "SZL"
	THB FiatCurrency = "THB"
	TJS FiatCurrency = "TJS"
	TMT FiatCurrency = "TMT"
	TND FiatCurrency = "TND"
	TOP FiatCurrency = "TOP"
	TRY FiatCurrency = "TRY"
	TTD FiatCurrency = "TTD"
	TWD FiatCurrency = "TWD"
	TZS FiatCurrency = "TZS"
	UAH FiatCurrency = "UAH"
	UGX FiatCurrency = "UGX"
	USD FiatCurrency = "USD"
	UYU FiatCurrency = "UYU"
	UZS FiatCurrency = "UZS"
	VEF FiatCurrency = "VEF"
	VES FiatCurrency = "VES"
	VND FiatCurrency = "VND"
	VUV FiatCurrency = "VUV"
	WST FiatCurrency = "WST"
	XAF FiatCurrency = "XAF"
	XAG FiatCurrency = "XAG"
	XAU FiatCurrency = "XAU"
	XCD FiatCurrency = "XCD"
	XDR FiatCurrency = "XDR"
	XOF FiatCurrency = "XOF"
	XPD FiatCurrency = "XPD"
	XPF FiatCurrency = "XPF"
	XPT FiatCurrency = "XPT"
	YER FiatCurrency = "YER"
	ZAR FiatCurrency = "ZAR"
	ZMK FiatCurrency = "ZMK"
	ZMW FiatCurrency = "ZMW"
	ZWL FiatCurrency = "ZWL"
)

type CurrencyPair string
const (
	BAT_USDC  CurrencyPair = "BAT-USDC"
	BTC_EUR   CurrencyPair = "BTC-EUR"
	BTC_GBP   CurrencyPair = "BTC-GBP"
	BTC_USD   CurrencyPair = "BTC-USD"
	ETH_BTC   CurrencyPair = "ETH-BTC"
	LINK_USDC CurrencyPair = "LINK-USDC"
)