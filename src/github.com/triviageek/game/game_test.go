package game

import (
	"testing"
	"time"
)

func TestCreateGame(t *testing.T) {
	Init()
	questions, game := JoinAGame()
	if len(game.toPlayers) == 0 {
		t.Fatal("A game cannot be created with no players")
	}
	if game.StartTime.Before(time.Now()) {
		t.Fatal("Game should not have started yet")
	}
	q := <-questions
	if game.StartTime.After(time.Now()) {
		t.Fatal("Game should have started now")
	}
	if q.Smell.Description == "" {
		t.Fatal("question.smell is null")
	}
	if q.Suggestions == nil {
		t.Fatal("question.suggestions is null")
	}
	if len(q.Suggestions) != 3 {
		t.Fatal("question does not contains 3 proposals", len(q.Suggestions))
	}
	if len(q.Suggestions[0]) == 0 {
		t.Fatal("suggestion should not be empty", q)
	}
}
