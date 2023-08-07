package controllers

import (
	"PY1/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func IndexRoute(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello skrrr!")
	if err != nil {
		return
	}
}
func Parse(w http.ResponseWriter, r *http.Request) {
	// newCode is responsible to save the given input
	var newCode models.SourceCode
	// consoleResponse is responsible of returning all of the console logs
	var consoleResponse models.ConsoleResponse
	// getting the body from the request
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "ERROR")
	}

	// parsing the json
	json.Unmarshal(reqBody, &newCode)
	// printing the input
	fmt.Println(newCode.Code)
	// TODO just setting a response
	consoleResponse.Console = "respuesta"
	json.NewEncoder(w).Encode(consoleResponse)
}
