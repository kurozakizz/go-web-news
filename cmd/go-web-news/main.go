package main

import (
	"net/http"
	"github.com/kurozakizz/go-web-news/pkg/app"
)

func main() {
	mux := http.NewServeMux()
	app.Mount(mux)
}