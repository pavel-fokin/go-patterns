package main

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExecuteTasks(t *testing.T) {
	ctx := context.Background()

	t.Run("No tasks", func(t *testing.T) {
		err := ExecuteTasks(ctx, []TaskFunc{}, 0)
		assert.NoError(t, err)
	})

	t.Run("All tasks succeed", func(t *testing.T) {
		tasks := []TaskFunc{
			func(ctx context.Context) error {
				return nil
			},
			func(ctx context.Context) error {
				return nil
			},
			func(ctx context.Context) error {
				return nil
			},
		}

		err := ExecuteTasks(ctx, tasks, len(tasks))
		assert.NoError(t, err)
	})

	t.Run("All tasks failed", func(t *testing.T) {
		expectedErr := errors.New("task failed")

		tasks := []TaskFunc{
			func(ctx context.Context) error {
				return expectedErr
			},
			func(ctx context.Context) error {
				return expectedErr
			},
			func(ctx context.Context) error {
				return expectedErr
			},
		}

		err := ExecuteTasks(ctx, tasks, len(tasks))
		assert.ErrorIs(t, err, expectedErr)
		assert.ErrorContains(
			t, err, "failed to execute concurrently 3 tasks, first error: task failed",
		)
	})

	t.Run("One task fails", func(t *testing.T) {
		expectedErr := errors.New("task failed")

		tasks := []TaskFunc{
			func(ctx context.Context) error {
				return nil
			},
			func(ctx context.Context) error {
				return expectedErr
			},
			func(ctx context.Context) error {
				return nil
			},
		}

		err := ExecuteTasks(ctx, tasks, len(tasks))
		assert.ErrorIs(t, err, expectedErr)
		assert.ErrorContains(
			t, err, "failed to execute concurrently 3 tasks, first error: task failed",
		)
	})

	t.Run("Context already canceled", func(t *testing.T) {
		expectedErr := context.Canceled

		tasks := []TaskFunc{
			func(ctx context.Context) error {
				return nil
			},
			func(ctx context.Context) error {
				return nil
			},
			func(ctx context.Context) error {
				return nil
			},
		}

		cancelCtx, cancel := context.WithCancel(ctx)
		cancel()

		err := ExecuteTasks(cancelCtx, tasks, len(tasks))
		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("One task is canceled", func(t *testing.T) {
		expectedErr := context.Canceled

		tasks := []TaskFunc{
			func(ctx context.Context) error {
				return sleepWithCtx(ctx, time.Second)
			},
			func(ctx context.Context) error {
				return nil
			},
			func(ctx context.Context) error {
				return nil
			},
		}

		cancelCtx, cancel := context.WithCancel(ctx)
		go func() {
			time.Sleep(time.Millisecond)
			cancel()
		}()

		err := ExecuteTasks(cancelCtx, tasks, len(tasks))
		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("All tasks are canceled", func(t *testing.T) {
		expectedErr := context.Canceled

		tasks := []TaskFunc{
			func(ctx context.Context) error {
				return sleepWithCtx(ctx, time.Second)
			},
			func(ctx context.Context) error {
				return sleepWithCtx(ctx, 2*time.Second)
			},
			func(ctx context.Context) error {
				return sleepWithCtx(ctx, 2*time.Second)
			},
		}

		cancelCtx, cancel := context.WithCancel(ctx)
		go func() {
			time.Sleep(time.Millisecond)
			cancel()
		}()

		err := ExecuteTasks(cancelCtx, tasks, len(tasks))
		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("One task fails others are canceled", func(t *testing.T) {
		expectedErr := errors.New("task failed")

		tasks := []TaskFunc{
			func(ctx context.Context) error {
				return sleepWithCtx(ctx, time.Second)
			},
			func(ctx context.Context) error {
				return expectedErr
			},
			func(ctx context.Context) error {
				return sleepWithCtx(ctx, time.Second)
			},
		}

		cancelCtx, cancel := context.WithCancel(ctx)
		go func() {
			time.Sleep(time.Millisecond)
			cancel()
		}()

		err := ExecuteTasks(cancelCtx, tasks, len(tasks))
		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("Tasks take longer than the context deadline", func(t *testing.T) {
		expectedErr := context.DeadlineExceeded

		tasks := []TaskFunc{
			func(ctx context.Context) error {
				return sleepWithCtx(ctx, time.Second)
			},
			func(ctx context.Context) error {
				return sleepWithCtx(ctx, time.Second)
			},
			func(ctx context.Context) error {
				return sleepWithCtx(ctx, time.Second)
			},
		}

		deadline := time.Now().Add(10 * time.Millisecond)
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()

		err := ExecuteTasks(ctx, tasks, len(tasks))
		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("Limit the number of workers", func(t *testing.T) {
		taskDuration := 10 * time.Millisecond

		// Create a list of tasks that sleep for 50 milliseconds.
		tasks := []TaskFunc{
			func(ctx context.Context) error {
				return sleepWithCtx(ctx, taskDuration)
			},
			func(ctx context.Context) error {
				return sleepWithCtx(ctx, taskDuration)
			},
			func(ctx context.Context) error {
				return sleepWithCtx(ctx, taskDuration)
			},
			func(ctx context.Context) error {
				return sleepWithCtx(ctx, taskDuration)
			},
		}

		// Measure the time it takes to execute all tasks.
		start := time.Now()
		err := ExecuteTasks(ctx, tasks, 2)
		elapsed := time.Since(start)

		assert.NoError(t, err)
		// The total time should be around 2 times the task duration.
		assert.InDelta(t, elapsed, 2*taskDuration, float64(taskDuration/2))
	})

	t.Run("Tasks are canceled with workers limit", func(t *testing.T) {
		expectedErr := context.Canceled

		tasks := []TaskFunc{
			func(ctx context.Context) error {
				return sleepWithCtx(ctx, time.Second)
			},
			func(ctx context.Context) error {
				return sleepWithCtx(ctx, time.Second)
			},
			func(ctx context.Context) error {
				return sleepWithCtx(ctx, time.Second)
			},
			func(ctx context.Context) error {
				return sleepWithCtx(ctx, time.Second)
			},
		}

		cancelCtx, cancel := context.WithCancel(ctx)
		go func() {
			time.Sleep(time.Millisecond)
			cancel()
		}()

		err := ExecuteTasks(cancelCtx, tasks, 2)
		assert.ErrorIs(t, err, expectedErr)
	})
}

// sleepWithCtx sleeps for the given duration or until the context is canceled.
func sleepWithCtx(ctx context.Context, duration time.Duration) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(duration):
		return nil
	}
}
