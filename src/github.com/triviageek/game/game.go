package game

import (
	"fmt"
	"time"
)

var runningGames []*Game

type Player struct {
	pseudo string
	game   *Game
	score  int
}

type Game struct {
	started bool
	ticker  *time.Ticker
	players []chan Question
}

func createOrJoinAGame() (<-chan Question, *Game) {
	qChan := make(chan Question, 1)
	for _, game := range runningGames {
		if game.started == false {
			game.players = append(game.players, qChan)
			return qChan, game
		}
	}
	newGame := &Game{ticker: time.NewTicker(20 * time.Second), players: []chan Question{qChan}}
	runningGames = append(runningGames, newGame)
	go newGame.start()
	return qChan, newGame
}

func (g *Game) start() {
	// Send a question every 20 sec
	for range g.ticker.C {
		g.started = true
		fmt.Println("Send questions to player(s)", len(g.players))
		q := <-Questions
		for _, player := range g.players {
			player <- q
		}
	}
}
