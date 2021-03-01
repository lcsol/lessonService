package repository

import "github.com/inspiritvr-organization/lesson-service-draft/pkg/entity"

// LessonRepository is an interface for any repository implementation
type LessonRepository interface {
	GetAll() ([]entity.Lesson, error)
	GetLessonByID(id string) (*entity.Lesson, error)
	CreateLesson(lesson entity.Lesson) (string, error)
	UpdateInfo(id string, lessonInfo entity.LessonInfo) (string, error)
	UpdateModelItem(id string, model entity.ModelItem) (int64, error)
	UpdateLabel(id string, label entity.Label) (int64, error)
	DeleteLesson(id string) (int64, error)
}
