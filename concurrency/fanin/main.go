package main

import "fmt"

func main() {
	Alice := Say(
		"Alice", "hi", "how'r you doing?", "i'm good",
	)
	Bob := Say(
		"Bob", "hey", "what's up", "great!",
	)

	all := FanIn(Alice, Bob)

	for i := 0; i < 6; i++ {
		fmt.Println(<-all)
	}
}
