package main

import "fmt"
import "math/rand"
import "time"

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newHand() deck {
	h := deck{}
	return h
}

func deal(d deck, h deck, handSize int) (deck, deck) {
	dmap := map2deck(d)

	for i := 0; i < handSize; i++ {
		rand.NewSource(time.Now().UnixNano())
		item := rand.Intn(len(d))
		h = append(h, dmap[item])
		delete(dmap, item)
	}

	d = getRemainingDeck(d, dmap)

	return d, h
}

func map2deck(d deck) map[int]string {
	dmap := map[int]string{}

	for i, card := range d {
		dmap[i] = card
	}

	return dmap
}

func getRemainingDeck(d deck, dmap map[int]string) deck {
	d = deck{}

	for _, card := range dmap {
		d = append(d, card)
	}

	return d
}
