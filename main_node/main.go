package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Du-bem/dominari_scriptum/main_node/db"
	"github.com/Du-bem/dominari_scriptum/main_node/server"
	"github.com/Du-bem/dominari_scriptum/main_node/types"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

func run(ctx context.Context, quit chan error) {
	fmt.Println("Running the Admin API with with privated Key!")

	var err error
	var dbName string
	var serverAPI types.ServerAPI
	if os.Getenv("USER") == "admin" {
		serverAPI.AccountWalletInfo, err = server.NewAdminAPI(ctx) // Admin API
		dbName = "admin.db"
	} else {
		serverAPI.AccountWalletInfo, err = server.NewUserAPI(ctx) // User API
		dbName = "user.db"
	}

	if err != nil {
		quit <- err
		return
	}

	// If an error occurs when initializing the db, shutdown immediately
	serverAPI.DBInfo, err = db.NewDatabase(ctx, dbName)
	if err != nil {
		quit <- err
		return
	}

	go server.RunServer(serverAPI) // Run the http server in a goroutine
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
