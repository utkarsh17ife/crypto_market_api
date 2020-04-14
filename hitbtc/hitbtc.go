package hitbtc

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"github.com/utkarsh17ife/crypto_market_api/store"
)

type HitBtc struct {
	store     *store.Store
	wsconn    *websocket.Conn
	configStr string
}

type params struct {
	Symbol string `json:"symbol"`
}

type payload struct {
	Method string `json:"method"`
	Params params `json:"params"`
	Id     int    `json:"id"`
}

func NewConnection(store *store.Store, host, channel string) *HitBtc {

	hb := &HitBtc{}
	hb.store = store

	u := url.URL{Scheme: "wss", Host: host, Path: channel}
	hb.configStr = u.String()

	hb.Connect()

	go hb.listen()

	return hb
}

func (hb *HitBtc) AddSymbol(symbol string) error {

	p := payload{
		Method: "subscribeTicker",
		Params: params{
			Symbol: symbol,
		},
		Id: 123,
	}

	err := hb.Write(p)
	if err != nil {
		return err
	}

	return nil
}

func (hb *HitBtc) Connect() *websocket.Conn {
	if hb.wsconn != nil {
		return hb.wsconn
	}
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for ; ; <-ticker.C {
		ws, _, err := websocket.DefaultDialer.Dial(hb.configStr, nil)
		if err != nil {
			continue
		}
		fmt.Println("hitBtc: Connect: Connected to HitBtc")
		hb.wsconn = ws
		return hb.wsconn
	}
}

func (hb *HitBtc) listen() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		for {
			ws := hb.Connect()
			if ws == nil {
				return
			}
			_, bytMsg, err := ws.ReadMessage()
			if err != nil {
				hb.Stop()
				break
			}
			hb.store.SaveTickerData(bytMsg)
		}
	}
}

func (hb *HitBtc) Write(payload interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	ws := hb.Connect()
	if ws == nil {
		err := fmt.Errorf("conn.ws is nil")
		return err
	}

	if err := ws.WriteMessage(
		websocket.TextMessage,
		data,
	); err != nil {
		fmt.Println("hitBtc: Write: Web socket write error")
	}
	fmt.Println("hitBtc: Write: Message sent successfully")
	return nil
}

func (hb *HitBtc) Stop() {
	if hb.wsconn != nil {
		hb.wsconn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		hb.wsconn.Close()
		hb.wsconn = nil
	}
}
