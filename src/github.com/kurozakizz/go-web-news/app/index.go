package app

import (
	"net/http"
	"github.com/kurozakizz/go-web-news/view"
)

func index(w http.ResponseWriter, r *http.Request) {
	view.Index(w, nil)
}
