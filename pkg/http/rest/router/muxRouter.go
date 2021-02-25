package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct {
	router *mux.Router
}

// NewMuxRouter returns a router implemented with mux
func NewMuxRouter() Router {
	return &muxRouter{mux.NewRouter()}
}

func (mr *muxRouter) GET(uri string, f func(rw http.ResponseWriter, r *http.Request)) {
	mr.router.HandleFunc(uri, f).Methods("GET")
}
