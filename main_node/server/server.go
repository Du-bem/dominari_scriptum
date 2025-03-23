package server

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Du-bem/dominari_scriptum/main_node/types"
	"github.com/btnguyen2k/consu/checksum"
	"github.com/gorilla/mux"
)

type serverAPI types.ServerAPI

// handleAccountBalance handles serving the account balance UI request
func (s serverAPI) handleAccountBalance(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" AccountBalance is Serving: ", r.URL.String())

	fmt.Fprintf(w, `{"balance": "%d"}`, s.GetBalance())
}

// handleAccountTx handles serving the account tx identified by the tx id provided
func (s serverAPI) handleAccountTx(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" AccountTx is Serving: ", r.URL.String())

	vars := mux.Vars(r)
	txID, ok := vars["key"]
	if ok {
		data, err := s.GetTransaction(txID)
		if err == nil {
			jData, _ := json.Marshal(data)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jData)
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, `{"error" : "%s"}`, err)
		}
	}
}

// handleAccountTxs handles serving the account txs associated with the current account
func (s serverAPI) handleAccountTxs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" AccountTxs is Serving: ", r.URL.String())

	data, err := s.GetTransactions()
	if err == nil {
		jData, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"error" : "%s"}`, err)
	}
}

// handleListSatelliteData handles serving the satellite data stored locally
func (s serverAPI) handleListSatelliteData(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" ListSatelliteData is Serving: ", r.URL.String())

	if adminMiddleware(w) {
		return // exit prematurely
	}

	vars := mux.Vars(r)
	val, ok := vars["offset"]
	if ok {
		offset, err := strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, `{"error" : "%s"}`, err)
			return
		}

		data, err := s.List(offset)
		if err == nil {
			jData, _ := json.Marshal(data)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jData)
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, `{"error" : "%s"}`, err)
		}
	}
}

// handleSearchSatelliteByID handles searching the locally stored satellite data by id
func (s *serverAPI) handleSearchSatelliteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" SearchSatelliteByID is Serving: ", r.URL.String())

	if adminMiddleware(w) {
		return // exit prematurely
	}

	vars := mux.Vars(r)
	txID, ok := vars["txid"]
	if ok {
		id, err := strconv.Atoi(txID)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, `{"error" : "%s"}`, err)
			return
		}

		data, err := s.SearchByID(id)
		if err == nil {
			jData, _ := json.Marshal(data)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jData)
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, `{"error" : "%s"}`, err)
		}
	}
}

// handleSatelliteDataStore generates a SHA256 checksum of the data before storing
// in the database.
func (s *serverAPI) handleSatelliteDataStore(w http.ResponseWriter, r *http.Request) {
	if adminMiddleware(w) {
		return // exit prematurely
	}

	var data types.SatelliteRequestData

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error" : "%s"}`, err)
		return
	}

	if data.Name == "" || len(data.Position) == 0 || len(data.Position) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error" : "empty data fields found"}`)
		return
	}

	checkSumStr := hex.EncodeToString(checksum.Sha256Checksum(data))
	_, err = s.Insert(checkSumStr, &data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error" : "%s"}`, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// adminMiddleware restricts access to end point only to the admin
func adminMiddleware(w http.ResponseWriter) bool {
	if os.Getenv("USER_TYPE") != "admin" {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"error": "missing permissions"}`)
		return true
	}
	return false
}

// RunServer takes a server API interface that exposes all the methods needed to
// serve all the http requests.
func RunServer(val types.ServerAPI) {
	s := serverAPI(val)
	router := mux.NewRouter()
	router.HandleFunc("/ui/balance", s.handleAccountBalance).Methods("GET")
	router.HandleFunc("/ui/tx/{key}", s.handleAccountTx).Methods("GET")
	router.HandleFunc("/ui/txs", s.handleAccountTxs).Methods("GET")

	// Protected endpoints: Only accessible to admin
	router.HandleFunc("/ui/list/{offset:[0-9]+}", s.handleListSatelliteData).Methods("GET")
	router.HandleFunc("/ui/search/{txid:[0-9]+}", s.handleSearchSatelliteByID).Methods("GET")
	router.HandleFunc("/data", s.handleSatelliteDataStore).Methods("POST")

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Running the Server on: ", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
