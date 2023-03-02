package main

import "fmt"

func main() {
	d := newDeck()
	h := newHand()
	d, h = deal(d, h, 5)
	h.print()
	fmt.Println("---------------")
	d.print()
}
