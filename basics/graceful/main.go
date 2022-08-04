package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	// Create a root context that can be canceled
	ctx, cancel := context.WithCancel(context.Background())

	// Goroutine listens for OS signal
	// and cancel context when recieve SIGINT or SIGTERM
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		// Wait for a signal
		<-sig
		// Cancel root context
		cancel()
	}()

	// Create an errgroup that will syncronize cancelation for goroutines
	// This errgroup uses root context
	g, gCtx := errgroup.WithContext(ctx)

	// Create a web server
	httpServer := &http.Server{
		Addr: ":8000",
	}

	// Start it in goroutine that is managed by errgroup
	g.Go(func() error {
		fmt.Println("Start server...")
		return httpServer.ListenAndServe()
	})

	// Wait for the root context will be canceled
	g.Go(func() error {
		<-gCtx.Done()
		fmt.Println("Stopping server...")
		return httpServer.Shutdown(context.Background())
	})

	// Wait until all goroutines in the error group have been finished
	if err := g.Wait(); err != nil {
		fmt.Printf("exit reason: %s \n", err)
	}
}
