package game

import (
	"testing"
)

func TestGenerateQuestions(t *testing.T) {
	Init()
	q := <-store
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

func BenchmarkGenerate10000RandomNums(b *testing.B) {
	for i := 0; i < b.N; i++ {
		shuffledKeys := make(chan int, 300)
		go generateRandomSeries(shuffledKeys)
		for i := 0; i < 10000; i++ {
			<-shuffledKeys
		}
	}
}

func BenchmarkGenerate1000Questions(b *testing.B) {
	Init()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 1000; i++ {
			<-store
		}
	}
}
