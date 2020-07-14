package main

import (
	"fmt"
	"os"
	"strings"
)

type card struct {
	suit  string
	value string
}

func (c card) toString() string {
	return fmt.Sprintf("%s of %s", c.value, c.suit)
}

func (c card) serialize() string {
	return fmt.Sprintf("%s:%s", c.suit, c.value)
}

func deserializeCard(s string) card {
	c := strings.Split(s, ":")
	if len(c) != 2 {
		fmt.Println("Incorrect card format:", s)
		os.Exit(1)
	}

	return card{suit: c[0], value: c[1]}
}
