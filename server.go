package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func serveHttp() error {
	p := os.Getenv("HTTP_PORT")

	r := mux.NewRouter()

	r.PathPrefix("/blocks").HandlerFunc(blockHandler).Methods("GET")
	fileServe(r, "/", "./www")

	srv := &http.Server{
		Handler:      handlers.CompressHandler(r),
		Addr:         fmt.Sprintf(":%s", p),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Started http server, serving at %s", p)
	return srv.ListenAndServe()
}

func blockHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp, _ := json.Marshal(blocks)
	w.Write(resp)
}

func fileServe(router *mux.Router, prefix string, directory string) {
	router.PathPrefix(prefix).Handler(http.StripPrefix(prefix, http.FileServer(http.Dir(directory)))).Methods("GET")
	http.Handle(prefix, router)
}
