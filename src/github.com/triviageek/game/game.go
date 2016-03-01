package game

import (
	"fmt"
	"time"
)

var runningGames []*Game

type Game struct {
	started bool
	ticker  *time.Ticker
	players []chan Question
}

type Question struct {
	step        int
	smell       Smell
	suggestions []string
}

type Response struct {
	step        int
	success     bool
	suggestions []string
}

func CreateOrJoinAGame() (<-chan Question, *Game) {
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
	var currentStep int
	// Send a question every 20 sec
	for range g.ticker.C {
		currentStep++
		g.started = true
		fmt.Println("Send questions to player(s)", len(g.players))
		q := <-store
		q.step = currentStep
		for _, player := range g.players {
			player <- q
		}
	}
}
