package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"

	"github.com/triviageek/game"
)

func main() {
	fmt.Println("Application starting...")
	game.Init()

	http.Handle("/", websocket.Handler(triviaHandler))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}

}

func triviaHandler(ws *websocket.Conn) {
	defer ws.Close()

	incomingReq := ws.Read()

}
