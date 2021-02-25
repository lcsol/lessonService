package router

import "net/http"

// Router matches incoming requests to their respective handler
type Router interface {
	GET(uri string, f func(rw http.ResponseWriter, r *http.Request))
}
