package main

import (
	"github.com/kurozakizz/go-web-news/app"
	"net/http"
)

const port = ":8080"

func main() {
	mux := http.NewServeMux()
	app.CreateRouter(mux)
	http.ListenAndServe(port, mux)
}
