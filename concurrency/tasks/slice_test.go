package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {
	// Create a new Slice
	s := NewSlice[int](0)

	// Append a value to the Slice
	s.Append(10)

	// Verify that the value was appended correctly
	assert.Len(t, s.slice, 1)
	assert.Equal(t, 10, s.slice[0])
}

func TestAppendConcurrent(t *testing.T) {
	// Number of goroutines to spawn
	numGoroutines := 100

	// Create a new Slice
	s := NewSlice[int](numGoroutines)

	// Create a WaitGroup to synchronize goroutines
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Spawn multiple goroutines to concurrently append values to the Slice
	for i := 0; i < numGoroutines; i++ {
		go func(value int) {
			defer wg.Done()
			s.Append(value)
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Verify that all values were appended correctly
	expectedSlice := make([]int, numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		expectedSlice[i] = i
	}
	assert.Len(t, s.slice, numGoroutines)
	assert.ElementsMatch(t, expectedSlice, s.slice)
}
