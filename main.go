package main

import (
	"fmt"
	"sync"
	"time"
)

//  go run main.go deck.go card.go
func main() {
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}

	cards := newDeck()
	cards.shuffle()

	game := []deck{}

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(time.Second * 2)
			mutex.Lock()

			hand, remainingCards := deal(cards, 4)
			hand.shuffle()

			cards = remainingCards
			game = append(game, hand)

			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	for _, hand := range game {
		hand.print()
		fmt.Println()
	}
}
