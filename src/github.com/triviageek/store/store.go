package store

import (
	"math/rand"
	"time"
)

var Questions = make(chan Question, 1000)

type Question struct {
	smell       Smell
	suggestions []string
}

type Smell struct {
	name        string
	description string
}

func Init() {

	rand.Seed(time.Now().UnixNano())

	shuffledKeys := make(chan int, 300)
	go generateRandomSeries(shuffledKeys)

	go func() {
		for {
			smell := smells[<-shuffledKeys]
			sugs := []string{smell.name, smells[<-shuffledKeys].name, smells[<-shuffledKeys].name}

			for i := range sugs {
				j := rand.Intn(i + 1)
				sugs[i], sugs[j] = sugs[j], sugs[i]
			}
			Questions <- Question{smell, sugs}
		}
	}()

}

func generateRandomSeries(c chan int) {
	for {
		keys := make(map[int]interface{}, len(smells))
		for i := 0; i < len(smells); i++ {
			keys[i] = struct{}{}
		}
		// Go runtime randomizes the iteration order access on map
		for k := range keys {
			c <- k
		}
	}
}
