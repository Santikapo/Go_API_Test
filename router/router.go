package routes

import (
	"net/http"
)

// server structure
type server struct {
	//db     *someDatabase
	router *http.ServeMux
}

// initalizes server routes
func (s *server) routes() {
	s.router.HandleFunc("", placeHolder)
	s.router.HandleFunc("", placeHolder)
	s.router.HandleFunc("", placeHolder)
}

// temp
func placeHolder(w http.ResponseWriter, r *http.Request) {
	return
}
