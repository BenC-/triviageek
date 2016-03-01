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

	receivedBytes := make([]byte, 100)
	n, err := ws.Read(receivedBytes)
	if err != nil {
		fmt.Println("Not a websocket request, close connection", err)
	}

	player := &game.Player{}
	if err := json.Unmarshal(receivedBytes[:n], player); err != nil {
		fmt.Println("Bad request, this is not a player. Close connection", err)
		return
	}

	fmt.Println("New player :", *player)

	questions, joiningGame := game.CreateOrJoinAGame()

	b, err := json.Marshal(joiningGame)
	if err != nil {
		fmt.Println("Unable to marshall game, WTF ?!", err)
		return
	}
	ws.Write(b)

	// Complete player
	player.Game = joiningGame

	// Send periodic questions
	go func(<-chan game.Question) {
		for question := range questions {
			fmt.Println(question)
			b, err := json.Marshal(question)
			if err != nil {
				fmt.Println("Unable to marshall question, WTF ?!", err)
				return
			}
			ws.Write(b)
		}
		fmt.Println("Game Ended")
		b, err := json.Marshal(player)
		if err != nil {
			fmt.Println("Unable to mashall player, WTF ?!", err)
			return
		}
		ws.Write(b)
		return
	}(questions)

	// Read responses
	for {
		n, err := ws.Read(receivedBytes)
		if err != nil {
			fmt.Println("Error while reading bytes", err)
		}
		fmt.Println(string(receivedBytes[:n]))
		response := &game.Response{}
		if err := json.Unmarshal(receivedBytes[:n], response); err != nil {
			fmt.Println("Bad request, this is not a response close connection", err)
			return
		}
		fmt.Println("Received response", response)
		if response.Success && response.Step >= joiningGame.Step {
			player.Score++
		}

	}

}
