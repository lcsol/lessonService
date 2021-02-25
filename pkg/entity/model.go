package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// A Model represents a model record in database
type Model struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" validate:"isdefault"`
	Name        string             `bson:"name,omitempty" validate:"required"`
	Description string             `bson:"description,omitempty"`
	Location    string             `bson:"location,omitempty" validate:"required"`
	Tag         []string           `bson:"tag,omitempty"`
}
