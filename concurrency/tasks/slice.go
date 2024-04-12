package main

import (
	"sync"
)

// Slice is a thread-safe slice.
type Slice[T any] struct {
	sync.Mutex
	slice []T
}

// NewSlice creates a new thread-safe slice.
func NewSlice[T any](size int) *Slice[T] {
	return &Slice[T]{slice: make([]T, 0, size)}
}

// Append method appends an item to the slice safely.
func (s *Slice[T]) Append(value T) {
	s.Lock()
	s.slice = append(s.slice, value)
	s.Unlock()
}

// Get method returns the item at the specified index.
func (s *Slice[T]) Get(index int) T {
	s.Lock()
	defer s.Unlock()
	return s.slice[index]
}

// Len method returns the length of the Slice.
func (s *Slice[T]) Len() int {
	s.Lock()
	defer s.Unlock()
	return len(s.slice)
}

// Slice method returns the underlying slice.
func (s *Slice[T]) Slice() []T {
	s.Lock()
	defer s.Unlock()
	return s.slice
}
