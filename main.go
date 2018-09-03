package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homesHandler).Methods("GET")
	fmt.Println("Basic Golang Api listening on port: 8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}
