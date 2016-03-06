package game

import (
	"math/rand"
	"time"
)

var store = make(chan Question, 1000)

type Smell struct {
	Name        string `json:"name, omitempty"`
	Description string `json:"description"`
}

func Init() {

	rand.Seed(time.Now().UnixNano())

	shuffledKeys := make(chan int, 300)
	go generateRandomSeries(shuffledKeys)

	go func() {
		for {
			smell := smells[<-shuffledKeys]
			sugs := []string{smell.Name, smells[<-shuffledKeys].Name, smells[<-shuffledKeys].Name}
			for i := range sugs {
				j := rand.Intn(i + 1)
				sugs[i], sugs[j] = sugs[j], sugs[i]
			}
			store <- Question{0, smell, sugs}
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
