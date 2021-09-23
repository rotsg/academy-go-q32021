package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rotsg/first_deliverable/usecase"
)

func GetField(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "response: invalid id")
		return
	}
	field, err := usecase.GetField(id)
	if err != nil {
		fmt.Fprintf(w, "response: %v\n", err)
	} else {
		json.NewEncoder(w).Encode(field)
	}
}
