package api

import (
	"encoding/json"
	"net/http"
)

type MercadoBitcoinAPI struct {
	BaseURL string
}

func NewMercadoBitcoinAPI() *MercadoBitcoinAPI {
	return &MercadoBitcoinAPI{
		BaseURL: "https://www.mercadobitcoin.net/api/",
	}
}

type Ticker struct {
	High string `json:"high"`
	Low  string `json:"low"`
	Vol  string `json:"vol"`
	Last string `json:"last"`
	Buy  string `json:"buy"`
	Sell string `json:"sell"`
	Date int64  `json:"date"`
	Open string `json:"open"`
}

func (mercadoBitcoinAPI *MercadoBitcoinAPI) GetTicker(asset string) (Ticker, error) {
	type Response struct {
		Ticker Ticker `json:"ticker"`
	}

	var response Response

	res, err := http.Get(mercadoBitcoinAPI.BaseURL + asset + "/ticker/")
	if err != nil {
		return response.Ticker, err
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return response.Ticker, err
	}

	return response.Ticker, nil
}

type Trade struct {
	Date   int64   `json:"date"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
	Price  float64 `json:"price"`
	TID    int64   `json:"tid"`
}

func (mercadoBitcoinAPI *MercadoBitcoinAPI) GetTrades(asset string) ([]Trade, error) {

	var trades []Trade

	res, err := http.Get(mercadoBitcoinAPI.BaseURL + asset + "/trades/")
	if err != nil {
		return trades, err
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&trades); err != nil {
		return trades, err
	}

	return trades, nil
}
