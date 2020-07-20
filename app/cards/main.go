/**
Go playground
- Building cards deck
*/
package main

import "fmt"

func main() {
	cards := retrieveData("cards.txt")
	cards.shuffleDeck()
	fmt.Println(cards)
}
