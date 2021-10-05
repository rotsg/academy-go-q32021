package main

import (
	"net/http"

	"github.com/rotsg/academy-go-q32021/route"
)

const port = ":8080"

func main() {
	http.ListenAndServe(port, route.New())
}
