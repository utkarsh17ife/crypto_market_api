package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/utkarsh17ife/crypto_market_api/hitbtc"
	"github.com/utkarsh17ife/crypto_market_api/routes"
	"github.com/utkarsh17ife/crypto_market_api/server"
	"github.com/utkarsh17ife/crypto_market_api/store"
)

func main() {

	// get access to store
	fmt.Println("main: Creating the store...")
	store := store.New()

	// passing the store to the websocket connection to receive and save the incoming data
	hitbtc := hitbtc.NewConnection(store, os.Getenv("HITBTC_HOST"), os.Getenv("HITBTC_CHANNEL"))

	symbols := os.Getenv("SYMBOLS")

	if symbols == "" {
		log.Fatal("main: Symbols missing from the configuration")
	}

	symbolsArray := strings.Split(symbols, "|")

	for _, symbol := range symbolsArray {
		hitbtc.AddSymbol(symbol)
	}

	server := server.New(store)

	fmt.Println("main: Initializing the routes...")
	routes := routes.InitRoutes(server)

	port := os.Getenv("HTTP_PORT")
	// gracefull server implementation pending
	fmt.Printf("main.go: Http server up on PORT: %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), routes))

}
