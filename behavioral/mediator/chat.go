package main

import "fmt"

type User struct {
	Name     string
	Messages []string
}

// Mediator
type Chat struct {
	Users map[string]*User
}

func NewChat() *Chat {
	return &Chat{make(map[string]*User)}
}

func (c *Chat) Add(user User) {
	c.Users[user.Name] = &user
}

func (c *Chat) Say(to User, msg string) error {
	user, ok := c.Users[to.Name]
	if !ok {
		return fmt.Errorf("%s not in the chat\n", to.Name)
	}
	user.Messages = append(user.Messages, msg)
	return nil
}

func (c *Chat) SayAll(msg string) {
	for _, user := range c.Users {
		user.Messages = append(user.Messages, msg)
	}
}
