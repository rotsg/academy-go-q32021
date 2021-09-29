package main

import (
	"net/http"

	"github.com/rotsg/bootcamp_challenge/route"
)

func main() {
	http.ListenAndServe(":8080", route.New())
}
