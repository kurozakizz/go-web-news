package app

import (
	"net/http"
	"github.com/kurozakizz/go-web-news/view"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	view.RenderIndexPage(w, nil)
}
