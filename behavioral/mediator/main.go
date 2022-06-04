package main

import (
	"fmt"
)

func main() {
	fmt.Println("Mediator pattern.")

	John := User{Name: "John"}
	Bob := User{Name: "Bob"}
	Alice := User{Name: "Alice"}

	chat := NewChat()
	chat.Add(Bob)
	chat.Add(John)

	chat.Say(Bob, "Hello, Bob!")
	chat.Say(Alice, "Hello, Alice!")
	chat.SayAll("Hello, All!")
}
