package getting

import (
	"github.com/inspiritvr-organization/lesson-service-draft/pkg/entity"
	repo "github.com/inspiritvr-organization/lesson-service-draft/pkg/repository"
)

// Service provides access to the lesson repository
type Service interface {
	FindAll() ([]entity.Lesson, error)
	Get(id string) (*entity.Lesson, error)
}

type service struct {
	lr repo.LessonRepository
}

// NewService creates a getLesson service with dependencies
func NewService(lr repo.LessonRepository) Service {
	return &service{lr}
}

func (s *service) FindAll() ([]entity.Lesson, error) {
	return s.lr.GetAll()
}

func (s *service) Get(id string) (*entity.Lesson, error) {
	return s.lr.GetLessonByID(id)
}
