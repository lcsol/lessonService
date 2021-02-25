package repository

import (
	"context"
	"fmt"

	"github.com/lcsol/lessonService/pkg/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// LessonCollection implements LessonRepository and represents a mongodb session with a lesson data model
type lessonCollection struct {
	coll *mongo.Collection
}

// NewLessonCollection creates a new LessonCollection
func NewLessonCollection(coll *mongo.Collection) LessonRepository {
	return &lessonCollection{coll}
}

// GetAll retrieves all labs data
func (l *lessonCollection) GetAll() ([]entity.Lesson, error) {
	ctx := context.TODO()
	list := []entity.Lesson{}

	cursor, err := l.coll.Find(ctx, bson.M{})
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
func (l *lessonCollection) GetLessonByID(id string) (*entity.Lesson, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	ctx := context.TODO()
	var lesson entity.Lesson
	err = l.coll.FindOne(ctx, bson.M{"_id": ID}).Decode(&lesson)
	if err != nil {
		return nil, err
	}
	return &lesson, err
}

// CreateLesson inserts a lesson into database
func (l *lessonCollection) CreateLesson(lesson entity.Lesson) (string, error) {
	res, err := l.coll.InsertOne(context.TODO(), lesson)
	oid := res.InsertedID.(primitive.ObjectID).Hex()
	return oid, err
}

// UpdateInfo updates the name, description, or tag of a lesson
func (l *lessonCollection) UpdateInfo(id string, lessonInfo entity.LessonInfo) (string, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}
	update := bson.D{{"$set",
		bson.D{
			{"name", lessonInfo.Name},
			{"description", lessonInfo.Description},
			{"tag", lessonInfo.Tag},
		},
	}}
	var updatedDoc bson.M
	err = l.coll.FindOneAndUpdate(context.TODO(), bson.M{"_id": ID}, update).Decode(&updatedDoc)
	if err != nil {
		return "", err
	}
	res := fmt.Sprintf("%v", updatedDoc[id])
	return res, err
}

// UpdateModelItem add a model to a lesson
func (l *lessonCollection) UpdateModelItem(id string, model entity.ModelItem) (int64, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}
	update := bson.M{"$push": bson.M{"models": bson.M{
		"modelId":  model.ModelID,
		"position": model.Position,
	}}}
	updateResult, err := l.coll.UpdateOne(context.TODO(), bson.M{"_id": ID}, update)
	if err != nil {
		return 0, err
	}
	return updateResult.ModifiedCount, err
}

// UpdateLabel add a label to a lesson
func (l *lessonCollection) UpdateLabel(id string, label entity.Label) (int64, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}
	update := bson.M{"$push": bson.M{"labels": bson.M{
		"content":  label.Content,
		"position": label.Position,
	}}}
	updateResult, err := l.coll.UpdateOne(context.TODO(), bson.M{"_id": ID}, update)
	if err != nil {
		return 0, err
	}
	return updateResult.ModifiedCount, err
}

// DeleteLesson deletes a lesson from database
func (l *lessonCollection) DeleteLesson(id string) (int64, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}
	res, err := l.coll.DeleteOne(context.TODO(), bson.M{"_id": ID})
	return res.DeletedCount, err
}
