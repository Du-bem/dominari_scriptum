package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func handleUIRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UI API data handler is not yet implemented!")
}

func handleSatelliteRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Satellite API data handler is not yet implemented!")
}

func RunServer() {
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
