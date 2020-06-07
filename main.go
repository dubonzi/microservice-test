package main

import (
	"log"
	"net/http"
)

const port = ":8080"

func main() {

	server := http.Server{Addr: port}

	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/thanks", ThanksHandler)

	log.Printf("Running on port %s", port)
	log.Println("[ERROR] server.ListenAndServe: ", server.ListenAndServe())
}

// RootHandler is the handler for the root route.
func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from AWS!"))
}

// ThanksHandler is the handler for the thanks route.
func ThanksHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Thank you Clara! :)"))
}
