package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"

	wallet "github.com/bitcoin-sv/spv-wallet-go-client"
	"github.com/bitcoin-sv/spv-wallet-go-client/config"
)

type AdminData struct {
	adminAPI *wallet.AdminAPI
}

func handleUIRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UI API data handler is not yet implemented!")
}

func handleSatelliteRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Satellite API data handler is not yet implemented!")
}

func (a *AdminData) server() {
	router := mux.NewRouter()
	router.HandleFunc("/ui", handleUIRequest).Methods("GET")
	router.HandleFunc("/data", handleSatelliteRequest).Methods("GET")

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Running the Server on: ", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}

func run(_ context.Context, quit chan error) {
	fmt.Println("Running the Admin API with with privated Key!")
	adminAPIData, err := wallet.NewAdminAPIWithXPriv(
		config.New(config.WithAddr(os.Getenv("WALLET_URL"))), os.Getenv("ADMIN_XPRIV"),
	)
	if err != nil {
		quit <- err
		return
	}

	admin := &AdminData{adminAPI: adminAPIData}

	go admin.server() // Run the http server in a goroutine
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
}
