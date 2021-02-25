package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ModelCollection represent a mongodb session with a model data
type ModelCollection struct {
	coll *mongo.Collection
}

// NewModelCollection creates a new ModelCollection
func NewModelCollection(coll *mongo.Collection) *ModelCollection {
	return &ModelCollection{coll}
}

// GetModelByID retrieves a single model by id
func (models *ModelCollection) GetModelByID(id string) (*Model, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	ctx := context.TODO()
	var model Model
	err = models.coll.FindOne(ctx, bson.M{"_id": ID}).Decode(&model)
	if err != nil {
		return nil, err
	}
	return &model, err
}
