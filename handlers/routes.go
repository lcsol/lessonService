package handlers

import "github.com/gorilla/mux"

// Routes returns a router matching incoming requests to their respective handler
func (lh *LessonHandler) Routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/lessons", lh.All).Methods("GET")
	router.HandleFunc("/lessons/{id}", lh.GetLesson).Methods("GET")
	router.HandleFunc("/lessons", lh.Create).Methods("POST")
	router.HandleFunc("/lessons/info/{id}", lh.UpdateLessonInfo).Methods("PUT")
	router.HandleFunc("/lessons/{id}", lh.Delete).Methods("DELETE")
	return router
}
