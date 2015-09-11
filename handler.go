package xwp

import (
	"fmt"
	"net/http"
	"strings"
)

type Handler struct {
	DB *DB
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if strings.HasPrefix(r.URL.Path, "/assets/") {
		http.ServeFile(w, r, r.URL.Path[1:])
		return
	}
	switch r.URL.Path {

	case "/":
		h.serveHome(w, r)

	default:
		http.NotFound(w, r)
	}
}

func (h *Handler) serveHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Sweet Home")

}
