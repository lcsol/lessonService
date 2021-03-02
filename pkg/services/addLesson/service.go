package adding

import (
	"github.com/inspiritvr-organization/lesson-service-draft/pkg/entity"
	repo "github.com/inspiritvr-organization/lesson-service-draft/pkg/repository"
)

// Service provides access to the lesson repository
type Service interface {
	AddLesson(lesson entity.Lesson) (string, error)
}

type service struct {
	lr repo.LessonRepository
}

// NewService creates a getLesson service with dependencies
func NewService(lr repo.LessonRepository) Service {
	return &service{lr}
}

func (s *service) AddLesson(lesson entity.Lesson) (string, error) {
	return s.lr.CreateLesson(lesson)
}
