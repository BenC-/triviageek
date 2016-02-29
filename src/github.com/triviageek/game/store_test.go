package game

import (
	"testing"
)

func TestGenerateQuestions(t *testing.T) {
	Init()
	q := <-Questions
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
			<-Questions
		}
	}
}
