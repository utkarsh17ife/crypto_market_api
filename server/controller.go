package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) ApiInfo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Crypto market data api"))
}

// the struct of data store is different from get all currencies response, so creating a custom response struct
type allCurrenciesResponse struct {
	Currencies interface{} `json:"currencies"`
}

func (s *Server) GetAllCurrencies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	currenciesData := s.Store.GetAllSymbols()

	response := allCurrenciesResponse{
		Currencies: currenciesData,
	}

	responseData, _ := json.Marshal(response)

	w.Write(responseData)

	return

}

func (s *Server) GetCurrencyBySymbol(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	symbol := params["symbol"]

	currencyData := s.Store.GetBySymbol(symbol)

	responseData, _ := json.Marshal(currencyData)

	w.Write(responseData)

	return

}
