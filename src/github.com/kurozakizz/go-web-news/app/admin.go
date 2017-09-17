package app

import (
	"github.com/kurozakizz/go-web-news/view"
	"net/http"
)

func adminLogin(w http.ResponseWriter, r *http.Request) {
	view.AdminLogin(w, nil)
}

func adminList(w http.ResponseWriter, r *http.Request) {

}

func adminCreate(w http.ResponseWriter, r *http.Request) {

}

func adminEdit(w http.ResponseWriter, r *http.Request) {

}
