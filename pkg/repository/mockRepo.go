package repository

import (
	"github.com/lcsol/lessonService/pkg/entity"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) GetAll() ([]entity.Lesson, error) {
	args := mock.Called()
	res := args.Get(0)
	return res.([]entity.Lesson), args.Error(1)
}

func (mock *MockRepository) GetLessonByID(id string) (*entity.Lesson, error) {
	args := mock.Called()
	res := args.Get(0)
	return res.(*entity.Lesson), args.Error(1)
}

func (mock *MockRepository) CreateLesson(lesson entity.Lesson) (string, error) {
	return "", nil
}
func (mock *MockRepository) UpdateInfo(id string, lessonInfo entity.LessonInfo) (string, error) {
	return "", nil
}
func (mock *MockRepository) UpdateModelItem(id string, model entity.ModelItem) (int64, error) {
	return 0, nil
}
func (mock *MockRepository) UpdateLabel(id string, label entity.Label) (int64, error) {
	return 0, nil
}
func (mock *MockRepository) DeleteLesson(id string) (int64, error) {
	return 0, nil
}
