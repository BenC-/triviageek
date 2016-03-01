package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"

	"github.com/triviageek/game"
)

func main() {
	fmt.Println("Application starting...")
	game.Init()

	http.Handle("/", websocket.Handler(triviaHandler))

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}

}

func triviaHandler(ws *websocket.Conn) {
	defer ws.Close()

	fmt.Println("Incoming connection from", ws.RemoteAddr().Network())

	frameReader, err := ws.NewFrameReader()

	if err != nil {
		fmt.Println("Not a websocket request, close connection", err)
		return
	}

	frameReader.TrailerReader()
	newPlayer := &game.Player{}
	if err := json.Unm(frameReader).Decode(newPlayer); err != nil {
		fmt.Println("Bad request, this is not a player. Close connection", err)
		return
	}

	questions, joiningGame := game.CreateOrJoinAGame()

	b, err := json.Marshal(joiningGame)
	if err != nil {
		fmt.Println("Unable to mashall new player, WTF ?!", err)
		return
	}

	// Write game
	ws.Write(b)

	// Send periodic questions
	go func(<-chan game.Question) {
		for question := range questions {
			b, err := json.Marshal(question)
			if err != nil {
				fmt.Println("Unable to mashall question, WTF ?!", err)
				return
			}
			ws.Write(b)
		}
	}(questions)

	// Read responses
	for {
		response := &game.Response{}
		if err := json.NewDecoder(frameReader).Decode(response); err != nil {
			fmt.Println("Bad request, this is not a questions close connection", err)
			return
		}
		fmt.Println("Received response", response)
	}

}
