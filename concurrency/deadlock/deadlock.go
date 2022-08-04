package main

import (
	"fmt"
	"sync"
)

var (
	lock1, lock2 sync.Mutex
)

func func1() {
	for {
		lock1.Lock()
		lock2.Lock()
		lock1.Unlock()
		lock2.Unlock()
	}
}

func func2() {
	for {
		// NOTICE: another order for locks
		lock2.Lock()
		lock1.Lock()
		lock1.Unlock()
		lock2.Unlock()
	}
}

func main() {
	ch := make(chan int)

	go func() {
		ch <- 1
	}()
	go func() {
		ch <- 2
	}()

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
