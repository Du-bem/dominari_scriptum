package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Du-bem/dominari_scriptum/main_node/types"
	"github.com/gorilla/mux"
)

type serverAPI types.ServerAPI

func (s serverAPI) handleAccBalance(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UI API data handler is not yet implemented!")
}
func (s serverAPI) handleAccTx(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UI API data handler is not yet implemented!")
}
func (s serverAPI) handleAccTxs(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UI API data handler is not yet implemented!")
}
func (s serverAPI) handleListSatData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UI API data handler is not yet implemented!")
}
func (s *serverAPI) handleSearchSatData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UI API data handler is not yet implemented!")
}

func (s *serverAPI) handleSatelliteData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Satellite API data handler is not yet implemented!")
}

func RunServer(val types.ServerAPI) {
	s := serverAPI(val)
	router := mux.NewRouter()
	router.HandleFunc("/ui/balance", s.handleAccBalance).Methods("GET")
	router.HandleFunc("/ui/tx", s.handleAccTx).Methods("GET")
	router.HandleFunc("/ui/txs", s.handleAccTxs).Methods("GET")
	router.HandleFunc("/ui/list", s.handleListSatData).Methods("GET")
	router.HandleFunc("/ui/search", s.handleSearchSatData).Methods("GET")

	router.HandleFunc("/data", s.handleSatelliteData).Methods("POST")

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Running the Server on: ", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
