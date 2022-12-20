package main

import "fmt"

func Say(who string, msgs ...string) <-chan string {
  c := make(chan string)
  go func() {
    for _, msg := range msgs {
      c <- fmt.Sprintf("%s said, %s", who, msg)
    }
    close(c)
  }()
  return c
}
