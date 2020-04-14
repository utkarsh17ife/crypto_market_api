package store

import (
	"encoding/json"
	"fmt"
	"sync"
)

type Currency struct {
	Ask         string `json:"ask"`
	Bid         string `json:"bid"`
	Last        string `json:"last"`
	Open        string `json:"open"`
	Low         string `json:"low"`
	High        string `json:"high"`
	Volume      string `json:"volume"`
	VolumeQuote string `json:"volumeQuote"`
	Timestamp   string `json:"timestamp"`
	Symbol      string `json:"symbol"`
}

type Currencies []Currency

type TickerResponse struct {
	Params Currency `params`
}

type Store struct {
	cryptoData map[string]Currency
	mutex      *sync.Mutex
}

func New() *Store {

	s := &Store{}
	s.cryptoData = make(map[string]Currency)
	s.mutex = &sync.Mutex{}

	return s
}

func (s *Store) SaveTickerData(marketData []byte) {

	tr := &TickerResponse{}

	err := json.Unmarshal(marketData, tr)
	if err != nil {
		fmt.Println("store: SaveTickerData: Cant process ticker data")
		return
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if tr.Params.Symbol == "" {
		return
	}

	s.cryptoData[tr.Params.Symbol] = tr.Params
}

func (s *Store) GetAllSymbols() Currencies {
	var currencies Currencies
	for _, currency := range s.cryptoData {
		currencies = append(currencies, currency)
	}
	return currencies
}

func (s *Store) GetBySymbol(symbol string) Currency {
	return s.cryptoData[symbol]
}
