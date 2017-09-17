package main

import (
	"net/http"
	"app"
)

const port = ":8080"

func main() {
	mux := http.NewServeMux()
	app.CreateRouter(mux)
	http.ListenAndServe(port, mux)
}
