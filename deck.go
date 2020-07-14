package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []card

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, card{suit, value})
		}
	}

	return cards
}

func deal(d deck, size int) (deck, deck) {
	return d[:size], d[size:]
}

func (d deck) shuffle() {
	// pseudorandom
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)

	for i := range d {
		random := generator.Intn(len(d) - 1)
		d[i], d[random] = d[random], d[i]
	}
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card.toString())
	}
}

func (d deck) serialize() string {
	cards := []string{}
	for _, card := range d {
		cards = append(cards, card.serialize())
	}
	return strings.Join(cards, ",")
}

func deserialize(value string) deck {
	cards := deck{}
	items := strings.Split(value, ",")
	for _, item := range items {
		cards = append(cards, deserializeCard(item))
	}
	return cards
}

func (d deck) saveToFile(filename string) error {
	buffer := []byte(d.serialize())
	err := ioutil.WriteFile(filename, buffer, 0666)

	return err
}

func newDeckFromFile(filename string) deck {
	buffer, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	value := string(buffer)
	return deserialize(value)
}
