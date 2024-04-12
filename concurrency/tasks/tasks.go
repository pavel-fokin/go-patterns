package main

import (
	"context"
	"errors"
	"fmt"
)

// Task represents an interface for executing tasks.
type Task interface {
	Execute(ctx context.Context) error
}

// TaskFunc represents a function that can be executed as a task.
type TaskFunc func(ctx context.Context) error

func (tf TaskFunc) Execute(ctx context.Context) error {
	return tf(ctx)
}

// ExecuteTasks executes a list of tasks concurrently.
// It takes a context, a list of tasks, and the number of workers to execute the tasks.
// It cancels remaining tasks if any of the tasks fail and returns an error.
func ExecuteTasks[T Task](ctx context.Context, tasks []T, workers int) error {
	if len(tasks) == 0 || workers == 0 {
		return nil
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Revise the number of workers.
	workers = min(workers, len(tasks))

	// Create a semaphore to limit the number of concurrent workers.
	semaphore := make(chan struct{}, workers)
	for i := 0; i < workers; i++ {
		semaphore <- struct{}{}
	}

	// Create a channel to collect errors from tasks.
	errCh := make(chan error, workers)

	// Define a worker function that executes a task.
	workerFunc := func(t Task) bool {
		select {
		case <-ctx.Done():
			return false
		case <-semaphore:
			go func() {
				// Release the semaphore when the task is done.
				defer func() {
					semaphore <- struct{}{}
				}()

				if err := t.Execute(ctx); err != nil {
					// Cancel the context if any task fails.
					cancel()
					if !errors.Is(err, context.Canceled) {
						errCh <- err
					}
				}
			}()
			return true
		}
	}

	// Execute each task concurrently.
	for _, task := range tasks {
		if !workerFunc(task) {
			break
		}
	}

	// Drain the semaphore.
	for i := 0; i < workers; i++ {
		<-semaphore
	}
	close(errCh)
	close(semaphore)

	// Return the first error that occurs.
	nErrors := len(errCh)
	if nErrors > 0 {
		err := <-errCh
		for range errCh {
			// Drain the error channel to make it garbage collected.
		}
		return fmt.Errorf(
			"failed to execute concurrently %d tasks, first error: %w",
			len(tasks),
			err,
		)
	}

	if ctx.Err() != nil {
		return ctx.Err()
	}

	return nil
}
