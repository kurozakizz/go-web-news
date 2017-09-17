package app

import (
	"net/http"
	"view"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	view.RenderIndexPage(w, nil)
}
