package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	getting "github.com/inspiritvr-organization/lesson-service-draft/pkg/services/getLesson"

	"github.com/gorilla/mux"
)

// LessonHandler is a handler to handle CRUD requests for lessons
type LessonHandler interface {
	All(rw http.ResponseWriter, r *http.Request)
	GetLesson(rw http.ResponseWriter, r *http.Request)
	serverError(rw http.ResponseWriter, err error)
}

type handler struct {
	infoLog *log.Logger
	errLog  *log.Logger
	g       getting.Service
}

// NewLessonHandler creates a new lesson handler
func NewLessonHandler(info *log.Logger, err *log.Logger, g getting.Service) LessonHandler {
	return &handler{info, err, g}
}

// All retrives all lessons info from database
func (h *handler) All(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")
	list, err := h.g.FindAll()
	if err != nil {
		h.serverError(rw, err)
	}
	err = json.NewEncoder(rw).Encode(list)
	if err != nil {
		h.serverError(rw, err)
	}
}

// GetLesson retrives a lesson by id
func (h *handler) GetLesson(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	lesson, err := h.g.Get(id)
	if err != nil {
		h.serverError(rw, err)
	}
	err = json.NewEncoder(rw).Encode(lesson)
	if err != nil {
		h.serverError(rw, err)
	}
}
