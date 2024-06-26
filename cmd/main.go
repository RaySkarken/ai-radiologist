package main

import (
	"ai-radiologist/internal/app"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	// Create a context to handle graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())

	// Create a WaitGroup to keep track of running goroutines
	var wg sync.WaitGroup

	// Start the HTTP server
	wg.Add(1)
	go func() {
		err := app.Run(ctx, &wg)
		if err != nil {
			log.Println(err)
		}
	}()

	// Listen for termination signals
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	// Wait for termination signal
	<-signalCh

	// Start the graceful shutdown process
	fmt.Println("\nGracefully shutting down HTTP server...")

	// Cancel the context to signal the HTTP server to stop
	cancel()

	// Wait for the HTTP server to finish
	wg.Wait()

	fmt.Println("Shutdown complete.")
}
