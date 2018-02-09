package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", Home).Methods("GET")
	router.HandleFunc("/home", Home).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func Home(w http.ResponseWriter, r *http.Request) {
	test := []byte("Welcome home")
	w.Write(test)
}
