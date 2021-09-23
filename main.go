package main

import (
	"net/http"

	"github.com/rotsg/first_deliverable/route"
)

func main() {
	http.ListenAndServe(":8080", route.New())
}
