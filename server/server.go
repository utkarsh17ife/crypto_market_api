package server

import (
	"github.com/utkarsh17ife/crypto_market_api/store"
)

type Server struct {
	Store *store.Store
}

func New(store *store.Store) *Server {

	srv := &Server{}
	srv.Store = store

	return srv

}
