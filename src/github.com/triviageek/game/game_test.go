package game

import (
	"testing"
)

func TestCreateGame(t *testing.T) {
	Init()
	questions, game := CreateOrJoinAGame()
	if len(game.players) == 0 {
		t.Fatal("A game cannot be created with no players")
	}
	if game.started == true {
		t.Fatal("Game should not have started yet")
	}
	q := <-questions
	if game.started == false {
		t.Fatal("Game should have started now")
	}
	if q.smell.description == "" {
		t.Fatal("question.smell is null")
	}
	if q.suggestions == nil {
		t.Fatal("question.suggestions is null")
	}
	if len(q.suggestions) != 3 {
		t.Fatal("question does not contains 3 proposals", len(q.suggestions))
	}
	if len(q.suggestions[0]) == 0 {
		t.Fatal("suggestion should not be empty", q)
	}
}
