package data

import (
	"context"
	"errors"
	"task_manager/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type TaskService struct {
	collection *mongo.Collection
}

func NewTaskService(collection *mongo.Collection) *TaskService {
	return &TaskService{
		collection: collection,
	}
}

func (ts *TaskService) AddTask(task models.Task) (models.Task,error) {
	if task.DueDate.Before(time.Now()) {
		return models.Task{},errors.New("due date can't be in the past")
	}

	insertResult, err := ts.collection.InsertOne(context.Background(), task)

	if err != nil {
		return models.Task{}, err
	}
	
	task.ID = insertResult.InsertedID.(primitive.ObjectID)
	return task, nil
}

func (ts *TaskService) GetAllTasks() ([]models.Task,error) {
	
	cursor, err := ts.collection.Find(context.Background(), bson.M{})

	if err != nil {
		return nil, err
	}

	var tasks []models.Task
	if err = cursor.All(context.Background(), &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (ts *TaskService) GetTask(id string) (task models.Task, err error){
	
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Task{}, errors.New("invalid id")
	}

	err = ts.collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&task)
	if err != nil {
		return models.Task{}, errors.New("task not found")
	}

	return task, nil

}

func (ts *TaskService) UpdateTask(id string, updatedTask models.Task) (err error) {

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid id")
	}

	_, err = ts.collection.UpdateOne(context.Background(), bson.M{"_id": objectID}, bson.M{"$set": updatedTask})
	if err != nil {
		return errors.New("task not found")
	}

	return nil
}

func (ts *TaskService) DeleteTask(id string) (err error) {

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid id")
	}

	_, err = ts.collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	if err != nil {
		return errors.New("task not found")
	}

	return nil
}