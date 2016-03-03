package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"

	"github.com/triviageek/game"
	http_ "github.com/triviageek/http"
)

func main() {
	fmt.Println("Application starting...")

	game.Init()

	http.Handle("/", websocket.Handler(http_.GameHandler))

	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}

}
