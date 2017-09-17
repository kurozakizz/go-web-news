package app

import (
	"view"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	view.Index(w, nil)
}
