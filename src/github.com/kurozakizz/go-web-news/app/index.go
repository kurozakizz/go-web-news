package app

import (
	"github.com/kurozakizz/go-web-news/view"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	view.Index(w, nil)
}
