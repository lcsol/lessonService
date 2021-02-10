package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// A Lesson represents a lesson record in database
type Lesson struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" validate:"isdefault"`
	Name        string             `bson:"name,omitempty" validate:"required"`
	Description string             `bson:"description,omitempty"`
	CreatedOn   time.Time          `bson:"createdOn,omitempty"`
	CreatorID   primitive.ObjectID `bson:"creatorId,omitempty"`
	Models      []ModelItem        `bson:"models" validate:"required"`
	Labels      []Label            `bson:"labels,omitempty"`
	Questions   [][]Question       `bson:"questions,omitempty"`
	Tag         []string           `bson:"tag,omitempty"`
}

// A LessonInfo includes the name, description, and tag of a lesson
type LessonInfo struct {
	Name        string   `bson:"name,omitempty" validate:"required"`
	Description string   `bson:"description,omitempty"`
	Tag         []string `bson:"tag,omitempty"`
}

// A ModelItem includes the id and position of a model in lesson
type ModelItem struct {
	ModelID  primitive.ObjectID `bson:"modelId,omitempty" validate:"required"`
	Position position           `bson:"position"`
}

// A Label includes the content and position of a label in lesson
type Label struct {
	Content  string   `bson:"content" validate:"required"`
	Position position `bson:"position"`
}

type position struct {
	X float64 `bson:"x"`
	Y float64 `bson:"y"`
	Z float64 `bson:"z"`
}

// type questions struct {
// 	Level       int            `bson:"level"`
// 	QuestionSet []questionItem `bson:"questionSet"`
// }

// A Question includes the information of a question in lesson
type Question struct {
	Level            int      `bson:"level"`
	Title            string   `bson:"title,omitempty"`
	Description      string   `bson:"description"`
	Choices          []string `bson:"choices"`
	CorrectChoiceIdx int      `bson:"correctChoiceIdx"`
}

// A Model represents a model record in database
type Model struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" validate:"isdefault"`
	Name        string             `bson:"name,omitempty" validate:"required"`
	Description string             `bson:"description,omitempty"`
	Location    string             `bson:"location,omitempty" validate:"required"`
	Tag         []string           `bson:"tag,omitempty"`
}
