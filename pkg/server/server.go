package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var logger ilogger

func Run(sk []byte, sl ilogger) {
	logger = sl

	r := mux.NewRouter()
	r.HandleFunc("/login", login).Methods("POST", "OPTIONS")
	r.HandleFunc("/logout", logout).Methods("GET")
	r.HandleFunc("/status", status).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
