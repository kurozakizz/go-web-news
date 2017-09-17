package app

import (
	"net/http"
)

// BindRouting handler mux
func CreateRouter(mux *http.ServeMux) {
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/news/", newsHandler)

	adminMux := http.NewServeMux()
	adminMux.HandleFunc("/login", adminLoginHandler)
	adminMux.HandleFunc("/list", adminListHandler)
	adminMux.HandleFunc("/create", adminCreateHandler)
	adminMux.HandleFunc("/edit", adminEditHandler)

	mux.Handle("/admin/", http.StripPrefix("/admin", onlyAdmin(adminMux)))
}

func onlyAdmin(h http.Handler) http.Handler {
	return h
}
