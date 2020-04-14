package routes

import (
	"github.com/gorilla/mux"
	"github.com/utkarsh17ife/crypto_market_api/server"
)

func InitRoutes(s *server.Server) *mux.Router {
	//Init Router
	r := mux.NewRouter()

	// arrange routes
	r.HandleFunc("/api/v1/info", s.ApiInfo).Methods("GET")
	r.HandleFunc("/api/v1/currency/all", s.GetAllCurrencies).Methods("GET")
	r.HandleFunc("/api/v1/currency/{symbol}", s.GetCurrencyBySymbol).Methods("GET")
	return r

}
