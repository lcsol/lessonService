package router

import (
	"github.com/gorilla/mux"
	"github.com/inspiritvr-organization/lesson-service-draft/pkg/http/rest/handlers"
)

// Routes returns a router matching incoming requests to their respective handler
func Routes(lh handlers.LessonHandler) *mux.Router {
	r := NewMuxRouter()
	r.GET("/lessons", lh.All)
	r.GET("/lessons/{id}", lh.GetLesson)

	return r.ReturnRouter()
}
