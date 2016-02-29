package game

import "time"

var runningGames []*Game = make([]*Game, 100)

type Player struct {
	pseudo string
	game   Game
	score  int
}

type Game struct {
	started bool
	ticker  *time.Ticker
	players []<-chan Question
}

func createOrJoinAGame() <-chan Question {
	qChan := make(chan Question, 1)
	for _, game := range runningGames {
		if game.started == false {
			game.players = append(game.players, qChan)
			return qChan
		}
	}
	newGame := &Game{ticker: time.NewTicker(20 * time.Second), players: []<-chan Question{qChan}}
	runningGames = append(runningGames, newGame)
	return qChan
}
