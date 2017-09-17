package app

import (
	"net/http"

	"github.com/kurozakizz/go-web-news/view"
)

func adminLoginHandler(w http.ResponseWriter, r *http.Request) {
	view.RenderAdminLoginPage(w, nil)
}

func adminListHandler(w http.ResponseWriter, r *http.Request) {

}

func adminCreateHandler(w http.ResponseWriter, r *http.Request) {

}

func adminEditHandler(w http.ResponseWriter, r *http.Request) {

}
