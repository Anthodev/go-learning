package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}
}

func TestNewDeckFirstCard(t *testing.T) {
	d := newDeck()

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}
}

func TestNewDeckLastCard(t *testing.T) {
	d := newDeck()

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs, but got %v", d[len(d)-1])
	}
}

func TestNewHand(t *testing.T) {
	d := newDeck()
	h := d.newHand()

	if len(h) != 0 {
		t.Errorf("Expected hand length different of 0, but got %v", len(h))
	}
}

func TestDeal(t *testing.T) {
	d := newDeck()
	h := d.newHand()
	d, h = d.deal(h, 5)

	if len(h) != 5 {
		t.Errorf("Expected hand length of 5, but got %v", len(h))
	}

	if len(d) != 47 {
		t.Errorf("Expected deck length of 47, but got %v", len(d))
	}
}
