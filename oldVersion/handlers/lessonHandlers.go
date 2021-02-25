package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"lessonService/models"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

// LessonHandler is a handler to handle CRUD requests for lessons
type LessonHandler struct {
	infoLog *log.Logger
	errLog  *log.Logger
	lessons *models.LessonCollection
	models  *models.ModelCollection
}

// NewLessonHandler creates a new lesson handler
func NewLessonHandler(info *log.Logger, err *log.Logger, lessons *models.LessonCollection, models *models.ModelCollection) *LessonHandler {
	return &LessonHandler{info, err, lessons, models}
}

// All calls GetAll func from labs to retrive all labs info from database
func (lh *LessonHandler) All(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")
	list, err := lh.lessons.GetAll()
	if err != nil {
		lh.serverError(rw, err)
	}
	err = json.NewEncoder(rw).Encode(list)
	if err != nil {
		lh.serverError(rw, err)
	}
}

// GetLesson calls GetLessonByID func from lessons to retrive a lesson by id
func (lh *LessonHandler) GetLesson(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	lesson, err := lh.lessons.GetLessonByID(id)
	if err != nil {
		lh.serverError(rw, err)
	}
	err = json.NewEncoder(rw).Encode(lesson)
	if err != nil {
		lh.serverError(rw, err)
	}
}

// Create calls CreateLesson func from lessons to insert a new lesson into database
func (lh *LessonHandler) Create(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")
	var lesson models.Lesson
	err := json.NewDecoder(r.Body).Decode(&lesson)
	if err != nil {
		lh.serverError(rw, err)
	}
	lesson.CreatedOn = time.Now()
	// request validation
	validate := validator.New()
	err = validate.Struct(lesson)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		lh.serverError(rw, validationErrors)
	} else {
		insert, err := lh.lessons.CreateLesson(lesson)
		if err != nil {
			lh.serverError(rw, err)
		} else {
			lh.infoLog.Printf("Created a new lab: %s", insert)
		}
	}
}

// UpdateLessonInfo calls UpdateInfo func from lessons to update a lesson name, description, or tags
func (lh *LessonHandler) UpdateLessonInfo(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var lessonInfo models.LessonInfo
	err := json.NewDecoder(r.Body).Decode(&lessonInfo)
	if err != nil {
		lh.serverError(rw, err)
	}
	validate := validator.New()
	err = validate.Struct(lessonInfo)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		lh.serverError(rw, validationErrors)
	} else {
		updatedDoc, err := lh.lessons.UpdateInfo(id, lessonInfo)
		if err != nil {
			lh.serverError(rw, err)
		}
		lh.infoLog.Printf("%s has been updated", updatedDoc)
	}
}

// UpdateLessonModels calls UpdateModelItem func from lessons to add a model to a lesson
func (lh *LessonHandler) UpdateLessonModels(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var model models.ModelItem
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		lh.serverError(rw, err)
	}
	validate := validator.New()
	err = validate.Struct(model)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		lh.serverError(rw, validationErrors)
	} else {
		updateResult, err := lh.lessons.UpdateModelItem(id, model)
		if err != nil {
			lh.serverError(rw, err)
		}
		lh.infoLog.Printf("Updated %v documents.\n", updateResult.ModifiedCount)
	}
}

// UpdateLessonLabels calls UpdateLabel func from lessons to add a label to a lesson
func (lh *LessonHandler) UpdateLessonLabels(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var label models.Label
	err := json.NewDecoder(r.Body).Decode(&label)
	if err != nil {
		lh.serverError(rw, err)
	}
	validate := validator.New()
	err = validate.Struct(label)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		lh.serverError(rw, validationErrors)
	} else {
		updateResult, err := lh.lessons.UpdateLabel(id, label)
		if err != nil {
			lh.serverError(rw, err)
		}
		lh.infoLog.Printf("Updated %v documents.\n", updateResult.ModifiedCount)
	}
}

// Delete calls DeleteLesson func from lessons to delete a lesson in database
func (lh *LessonHandler) Delete(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")
	id := mux.Vars(r)["id"]
	deleteRes, err := lh.lessons.DeleteLesson(id)
	if err != nil {
		lh.serverError(rw, err)
	}
	lh.infoLog.Printf("Deleted %d lesson", deleteRes.DeletedCount)
}

// GetModel calls GetModelByID func from models to retrive a model by id
func (lh *LessonHandler) GetModel(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	model, err := lh.models.GetModelByID(id)
	if err != nil {
		lh.serverError(rw, err)
	}
	err = json.NewEncoder(rw).Encode(model)
	if err != nil {
		lh.serverError(rw, err)
	}
}
