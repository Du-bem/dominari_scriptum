package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Du-bem/dominari_scriptum/main_node/server"
	_ "github.com/joho/godotenv/autoload"
)

func run(ctx context.Context, quit chan error) {
	fmt.Println("Running the Admin API with with privated Key!")

	var err error
	if os.Getenv("USER") == "admin" {
		_, err = server.NewAdminAPI(ctx) // Admin API
	} else {
		_, err = server.NewUserAPI(ctx) // User API
	}

	if err != nil {
		quit <- err
		return
	}

	go server.RunServer() // Run the http server in a goroutine
}

func main() {
	quit := make(chan error)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// timer that ticks after 10 seconds.
	timer := time.NewTicker(time.Second * 10)

	// Main context that shutdown any background threads before quiting
	ctx, cancelFunc := context.WithCancel(context.Background())

	go func() {
	loop:
		for {
			select {
			case <-c:
				// Ctrl+C Detected
				break loop
			case err := <-quit:
				// Program quit execution.
				fmt.Println("ERROR: ", err)
				break loop
			case <-timer.C:
				continue
			}
		}

		// Cancel the context function
		cancelFunc()
	}()

	run(ctx, quit)

	<-ctx.Done() // Exit the goroutine once the context is cancelled

	fmt.Println("!!!!! System shutting down !!!!!")
}
