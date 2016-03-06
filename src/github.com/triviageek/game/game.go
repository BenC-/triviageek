package game

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	NumOfQuestionsPerGame = 12
	QuestionPeriod        = 20
)

var runningGames []*Game

type Game struct {
	Name       string    `json:"name"`
	StartTime  time.Time `json:"startTime"`
	Step       int       `json:"step"`
	ticker     *time.Ticker
	players    []*Player
	resultMask []bool
	stopChan   chan interface{}
}

type Result struct {
	Players []*Player `json:"players"`
}

type Question struct {
	Step        int      `json:"step"`
	Smell       Smell    `json:"smell"`
	Suggestions []string `json:"suggestions"`
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
			oq := obfuscateQuestion(q)
			for _, player := range g.players {
				player.CurrentQuestion = q
				player.marshalAndSend(oq)
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

func obfuscateQuestion(q Question) Question {
	osmell := Smell{
		Description: q.Smell.Description,
	}
	oq := Question{
		Step:        q.Step,
		Smell:       osmell,
		Suggestions: q.Suggestions,
	}
	return oq
}
