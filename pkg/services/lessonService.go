package services

import "github.com/inspiritvr-organization/lesson-service-draft/pkg/entity"

type LessonService interface {
	FindAll() ([]entity.Lesson, error)
	Get(id string) (*entity.Lesson, error)
	Create(lesson entity.Lesson) (string, error)
	UpdateInfo(id string, lessonInfo entity.LessonInfo) (string, error)
	UpdateModelItem(id string, model entity.ModelItem) (int, error)
	UpdateLabel(id string, label entity.Label) (int, error)
	Delete(id string) (int, error)
}
