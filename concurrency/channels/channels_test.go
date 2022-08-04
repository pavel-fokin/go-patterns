package main

import "testing"

func BenchmarkUnbufferedChannelInt(b *testing.B) {
	ch := make(chan int)
	go func() {
		for {
			<-ch
		}
	}()
	for i := 0; i < b.N; i++ {
		ch <- i
	}
}

func BenchmarkBufferedChannelInt(b *testing.B) {
	ch := make(chan int, 1)
	go func() {
		for {
			<-ch
		}
	}()
	for i := 0; i < b.N; i++ {
		ch <- i
	}
}
