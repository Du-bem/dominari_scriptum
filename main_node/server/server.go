package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	

	wallet "github.com/bitcoin-sv/spv-wallet-go-client"
	"github.com/bitcoin-sv/spv-wallet-go-client/config"
)

type AdminData struct {
	adminAPI *wallet.AdminAPI
}

func NewAdminAPI() (*AdminData, error) {
	adminAPIData, err := wallet.NewAdminAPIWithXPriv(
		config.New(config.WithAddr(os.Getenv("WALLET_URL"))), os.Getenv("ADMIN_XPRIV"),
	)
	if err != nil {
		return nil, err
	}

	return &AdminData{adminAPI: adminAPIData}, nil
}

func handleUIRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UI API data handler is not yet implemented!")
}

func handleSatelliteRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Satellite API data handler is not yet implemented!")
}

func (a *AdminData) RunServer() {
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
