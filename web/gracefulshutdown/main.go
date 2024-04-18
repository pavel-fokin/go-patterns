package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

const maxShutdownDuration = 30 * time.Second

func main() {
	// Listen for OS signals.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Create a web server.
	httpServer := &http.Server{
		Addr: ":8000",
	}

	// Start the server.
	go func() {
		log.Println("Start server...")
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v\n", err)
		}
	}()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), maxShutdownDuration)
	defer cancel()

	// Shutdown the server gracefully.
	log.Println("Shutdown server...")
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to shutdown the server gracefully: %v\n", err)
	}
}
