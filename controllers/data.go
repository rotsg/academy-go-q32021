package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rotsg/first_deliverable/usecases"
)

func GetField(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "error: %v\n", err)
	} else {
		field, err := usecases.GetField(id)
		if err != nil {
			fmt.Fprintf(w, "error: %v\n", err)
		} else {
			json.NewEncoder(w).Encode(field)
		}
	}

}
