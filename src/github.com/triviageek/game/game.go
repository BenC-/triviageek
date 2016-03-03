package game

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	NumOfQuestionsPerGame = 4
	QuestionPeriod        = 5
)

var runningGames []*Game

type Game struct {
	Name      string    `json:"name,omitempty"`
	StartTime time.Time `json:"startTime,omitempty"`
	Step      int       `json:"step,omitempty"`
	ticker    *time.Ticker
	players   []*Player
	stopChan  chan interface{}
}

type Result struct {
	Players []*Player `json:"players"`
}

type Question struct {
	Step        int      `json:"step,omitempty"`
	Smell       Smell    `json:"smell,omitempty"`
	Suggestions []string `json:"suggestions,omitempty"`
}

func newGame() *Game {
	newGame := &Game{Name: randName(), StartTime: time.Now().Add(time.Second * QuestionPeriod), ticker: time.NewTicker(QuestionPeriod * time.Second), players: []*Player{}}
	return newGame
}

func (g *Game) start() {
	// Send a question every 20 sec
	for {
		select {
		case <-g.ticker.C:
			g.Step++
			if g.Step > NumOfQuestionsPerGame { // End of game
				g.broadcastResultAndClose()
			}
			fmt.Println(fmt.Sprintf("Send questions to %d player(s)", len(g.players)))
			q := <-store
			q.Step = g.Step
			for _, player := range g.players {
				player.marshalAndSend(q)
			}
		case <-g.stopChan:
			return
		}
	}
}

func (g *Game) broadcastResultAndClose() {
	result := &Result{g.players}
	for _, player := range g.players {
		player.marshalAndSend(result)
		player.endGame()
	}
	g.stopChan <- struct{}{}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randName() string {
	b := make([]rune, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
