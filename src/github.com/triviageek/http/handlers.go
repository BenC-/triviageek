package http

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"

	"github.com/triviageek/game"
)

func GameHandler(ws *websocket.Conn) {
	defer func(ws *websocket.Conn) {
		if err := recover(); err != nil {
			ws.Close()
		}
	}(ws)

	fmt.Println("Incoming connection from", ws.RemoteAddr())

	receivedBytes := make([]byte, 100)
	n, err := ws.Read(receivedBytes)
	if err != nil {
		fmt.Println("Not a websocket request, close connection", err)
		panic(err)
	}
	p := &game.Player{}
	if err := json.Unmarshal(receivedBytes[:n], p); err != nil {
		fmt.Println("Bad request, object is not a player. Close connection", err)
		panic(err)
	}
	p.Ws = ws

	fmt.Println("New player joining the game :", *p)

	p.JoinAGame()
	p.HandleEvents()
}
