package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/inspiritvr-organization/lesson-service-draft/pkg/entity"
	adding "github.com/inspiritvr-organization/lesson-service-draft/pkg/services/addLesson"
	getting "github.com/inspiritvr-organization/lesson-service-draft/pkg/services/getLesson"

	"github.com/gorilla/mux"
)

// LessonHandler is a handler to handle CRUD requests for lessons
type LessonHandler interface {
	All(rw http.ResponseWriter, r *http.Request)
	GetLesson(rw http.ResponseWriter, r *http.Request)
	Create(rw http.ResponseWriter, r *http.Request)
	serverError(rw http.ResponseWriter, err error)
}

type handler struct {
	infoLog *log.Logger
	errLog  *log.Logger
	get     getting.Service
	add     adding.Service
}

// NewLessonHandler creates a new lesson handler
func NewLessonHandler(
	info *log.Logger,
	err *log.Logger,
	get getting.Service,
	add adding.Service) LessonHandler {
	return &handler{info, err, get, add}
}

// All retrives all lessons info from database
func (h *handler) All(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")
	list, err := h.get.FindAll()
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
	lesson, err := h.get.Get(id)
	if err != nil {
		h.serverError(rw, err)
	}
	err = json.NewEncoder(rw).Encode(lesson)
	if err != nil {
		h.serverError(rw, err)
	}
}

// Create calls CreateLesson func from lessons to insert a new lesson into database
func (h *handler) Create(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")
	var lesson entity.Lesson
	err := json.NewDecoder(r.Body).Decode(&lesson)
	if err != nil {
		h.serverError(rw, err)
	}
	lesson.CreatedOn = time.Now()
	// request validation
	validate := validator.New()
	err = validate.Struct(lesson)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		h.serverError(rw, validationErrors)
	} else {
		insert, err := h.add.AddLesson(lesson)
		if err != nil {
			h.serverError(rw, err)
		} else {
			h.infoLog.Printf("Created a new lab: %s", insert)
		}
	}
}
