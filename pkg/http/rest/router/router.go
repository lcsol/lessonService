package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Router matches incoming requests to their respective handler
type Router interface {
	ReturnRouter() *mux.Router
	GET(uri string, f func(rw http.ResponseWriter, r *http.Request))
	POST(uri string, f func(rw http.ResponseWriter, r *http.Request))
}
