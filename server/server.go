package server // import "github.com/Santikapo/Go_API_Test/server"

import (
	"net/http"
)

// server structure
type server struct {
	//db     *someDatabase
	router *http.ServeMux
}
