package app

import (
	"net/http"
	"view"
)

func index(w http.ResponseWriter, r *http.Request) {
	view.Index(w, nil)
}