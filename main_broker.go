package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"taxi-tracker-broker/controller"
	"taxi-tracker-broker/pubsubimp"
)

func main() {
	port := "3001"
	router := mux.NewRouter().StrictSlash(false)

	topicsHandler := pubsubimp.NewTopicHandler()
	connectionHandler := controller.NewConnectionCreator(topicsHandler)
	router.HandleFunc("/ws/{clientId}", connectionHandler.CreateConnection)
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Ok"))
	}).Methods("GET")

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	fmt.Println("Listening in port:", port)
	server.ListenAndServe()
}
