package game

import (
	"fmt"
	"time"
)

const NumOfQuestionsPerGame = 12

var runningGames []*Game

type Game struct {
	StartTime time.Time `json:"startTime,omitempty"`
	Step      int       `json:"step,omitempty"`
	ticker    *time.Ticker
	toPlayers []chan Question
}

type Question struct {
	Step        int      `json:"step,omitempty"`
	Smell       Smell    `json:"smell,omitempty"`
	Suggestions []string `json:"suggestions,omitempty"`
}

type Response struct {
	Step    int  `json:"step,omitempty"`
	Success bool `json:"success,omitempty"`
}

func CreateOrJoinAGame() (<-chan Question, *Game) {
	qChan := make(chan Question, 1)
	for _, game := range runningGames {
		if game.StartTime.After(time.Now()) {
			game.toPlayers = append(game.toPlayers, qChan)
			return qChan, game
		}
	}
	newGame := &Game{StartTime: time.Now().Add(time.Second * 20), ticker: time.NewTicker(20 * time.Second), toPlayers: []chan Question{qChan}}
	runningGames = append(runningGames, newGame)
	go newGame.start()
	return qChan, newGame
}

func (g *Game) start() {
	// Send a question every 20 sec
	for range g.ticker.C {
		g.Step++
		if g.Step > NumOfQuestionsPerGame {
			for _, player := range g.toPlayers {
				close(player)
			}
		}
		fmt.Println("Send questions to player(s)", len(g.toPlayers))
		q := <-store
		q.Step = g.Step
		for _, player := range g.toPlayers {
			player <- q
		}
	}
}
