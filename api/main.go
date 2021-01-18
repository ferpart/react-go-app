package main

import (
	"encoding/json"
	"os"
	"os/signal"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var routes = []struct {
	path string
	method string
	handler http.HandlerFunc
}{
	{
		path:		"/api",
		method: 	http.MethodGet,
		handler: 	getAPI,
	},
}

func main() {
	defaultPort := 5000

	router := mux.NewRouter()

	for _, r := range routes {
		router.HandleFunc(r.path, r.handler).Methods(r.method)
	}

	srv := &http.Server {
		Addr:		fmt.Sprintf(":%d", defaultPort),
		Handler:	router,
	}

	go func() {
		log.Infof("starting server on port %d...\n", defaultPort)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
}

type response struct {
	Status string `json:"status"`
}

func getAPI(w http.ResponseWriter, r *http.Request){
	elem := response{"success"}
	setHeaders(&w)

	json.NewEncoder(w).Encode(elem)
}

func setHeaders(w *http.ResponseWriter){
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}