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

	var newCode models.SourceCode
	var consoleResponse models.ConsoleResponse

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "ERROR")
	}
	json.Unmarshal(reqBody, &newCode)

	fmt.Println(newCode.Code)

	consoleResponse.Console = "respuesta"

	json.NewEncoder(w).Encode(consoleResponse)
}
