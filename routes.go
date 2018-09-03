package main

import "net/http"

func homesHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("I am up and running!"))
}
