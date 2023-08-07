package main

import (
	"PY1/controllers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello mundo")
	router := mux.NewRouter().StrictSlash(true)
	parserRoutes(router)
	log.Fatal(http.ListenAndServe(":3000", router))
}

func parserRoutes(router *mux.Router) {
	router.HandleFunc("/", controllers.IndexRoute).Methods("GET")

}
