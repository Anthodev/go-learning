package main

import "fmt"

func main() {
	d := newDeck()
	h := d.newHand()
	d, h = d.deal(h, 5)
	h.print()
	fmt.Println("---------------")
	d.print()
}
