package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type lessonRepository interface {
}

// LessonCollection represent a mongodb session with a lesson data model
type LessonCollection struct {
	coll *mongo.Collection
}

// NewLessonCollection creates a new LessonCollection
func NewLessonCollection(coll *mongo.Collection) *LessonCollection {
	return &LessonCollection{coll}
}

// GetAll retrieves all labs data
func (lessons *LessonCollection) GetAll() ([]Lesson, error) {
	ctx := context.TODO()
	list := []Lesson{}

	cursor, err := lessons.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &list)
	if err != nil {
		return nil, err
	}
	return list, err
}

// GetLessonByID retrieves a single lesson by id
func (lessons *LessonCollection) GetLessonByID(id string) (*Lesson, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	ctx := context.TODO()
	var lesson Lesson
	err = lessons.coll.FindOne(ctx, bson.M{"_id": ID}).Decode(&lesson)
	if err != nil {
		return nil, err
	}
	return &lesson, err
}

// CreateLesson inserts a lesson into database
func (lessons *LessonCollection) CreateLesson(lesson Lesson) (*mongo.InsertOneResult, error) {
	return lessons.coll.InsertOne(context.TODO(), lesson)
}

// UpdateInfo updates the name, description, or tag of a lesson
func (lessons *LessonCollection) UpdateInfo(id string, lessonInfo LessonInfo) (*bson.M, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	update := bson.D{{"$set",
		bson.D{
			{"name", lessonInfo.Name},
			{"description", lessonInfo.Description},
			{"tag", lessonInfo.Tag},
		},
	}}
	var updatedDoc bson.M
	err = lessons.coll.FindOneAndUpdate(context.TODO(), bson.M{"_id": ID}, update).Decode(&updatedDoc)
	if err != nil {
		return nil, err
	}
	return &updatedDoc, err
}

// UpdateModelItem add a model to a lesson
func (lessons *LessonCollection) UpdateModelItem(id string, model ModelItem) (*mongo.UpdateResult, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	update := bson.M{"$push": bson.M{"models": bson.M{
		"modelId":  model.ModelID,
		"position": model.Position,
	}}}
	updateResult, err := lessons.coll.UpdateOne(context.TODO(), bson.M{"_id": ID}, update)
	if err != nil {
		return nil, err
	}
	return updateResult, err
}

// UpdateLabel add a label to a lesson
func (lessons *LessonCollection) UpdateLabel(id string, label Label) (*mongo.UpdateResult, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	update := bson.M{"$push": bson.M{"labels": bson.M{
		"content":  label.Content,
		"position": label.Position,
	}}}
	updateResult, err := lessons.coll.UpdateOne(context.TODO(), bson.M{"_id": ID}, update)
	if err != nil {
		return nil, err
	}
	return updateResult, err
}

// DeleteLesson deletes a lesson from database
func (lessons *LessonCollection) DeleteLesson(id string) (*mongo.DeleteResult, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return lessons.coll.DeleteOne(context.TODO(), bson.M{"_id": ID})
}
