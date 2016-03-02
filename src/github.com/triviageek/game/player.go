package game

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
	"time"
)

type Player struct {
	Pseudo string          `json:"pseudo"`
	Score  int             `json:"score"`
	Ws     *websocket.Conn `json:"-"`
}

type Response struct {
	Step    int  `json:"step"`
	Success bool `json:"success"`
}

func (p *Player) JoinAGame() {
	for _, game := range runningGames {
		if game.StartTime.After(time.Now()) {
			game.players = append(game.players, p)
			p.marshalAndSend(game)
			return
		}
	}
	newGame := newGame()
	newGame.players = append(newGame.players, p)
	runningGames = append(runningGames, newGame)
	p.marshalAndSend(newGame)
	go newGame.start()
}

func (p *Player) HandleEvents() {
	receivedBytes := make([]byte, 100)
	for {
		n, err := p.Ws.Read(receivedBytes)
		if err != nil {
			fmt.Println("Error while reading bytes or game ended")
			return
		}
		fmt.Println(string(receivedBytes[:n]))
		response := &Response{}
		if err := json.Unmarshal(receivedBytes[:n], response); err != nil {
			fmt.Println("Bad request, this is not a response close connection", err)
			return
		}
		fmt.Println("Received response from player", p.Pseudo, response)
		if response.Success {
			p.Score++
		}
	}
}

func (p *Player) marshalAndSend(o interface{}) {
	b, err := json.Marshal(o)
	if err != nil {
		fmt.Println("Unable to marshall", err)
		return
	}
	p.Ws.Write(b)
}

func (p *Player) endGame() {
	err := p.Ws.Close()
	if err != nil {
		fmt.Println("Error while closing websocket", err)
	}
}
