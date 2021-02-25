package getting

import (
	"testing"

	"github.com/lcsol/lessonService/pkg/entity"
	repo "github.com/lcsol/lessonService/pkg/repository"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestFindAll(t *testing.T) {
	mockRepo := new(repo.MockRepository)

	id := primitive.NewObjectID()

	lesson := entity.Lesson{ID: id, Name: "mock lesson", Description: "a lesson for mock"}
	// Setup expectations
	mockRepo.On("GetAll").Return([]entity.Lesson{lesson}, nil)
	testService := NewService(mockRepo)
	res, err := testService.FindAll()

	// Mock Assertion: Behavioral
	mockRepo.AssertExpectations(t)

	// Data Assertion
	assert.Equal(t, id, res[0].ID)
	assert.Equal(t, "mock lesson", res[0].Name)
	assert.Equal(t, "a lesson for mock", res[0].Description)
	assert.Nil(t, err)
}
