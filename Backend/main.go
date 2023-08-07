package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func indexRoute(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello skrrr!")
	if err != nil {
		return
	}
}

func main() {
	fmt.Println("Hello mundo")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	log.Fatal(http.ListenAndServe(":3000", router))
}
