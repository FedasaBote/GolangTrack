package data

import (
	"context"
	"enhanced_task_manager/models"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)




type TaskService struct {
	taskCollection *mongo.Collection
}

func NewTaskService(taskCollection *mongo.Collection) *TaskService {
	return &TaskService{
		taskCollection: taskCollection,
	}
}

func (ts *TaskService) AddTask(task models.Task) (models.Task, error) {
	if task.DueDate.Before(time.Now()) {
		return models.Task{}, errors.New("due date can't be in the past")
	}

	insertOneResult, err := ts.taskCollection.InsertOne(context.Background(), task)

	if err != nil {
		return models.Task{}, err
	}

	task.ID = insertOneResult.InsertedID.(primitive.ObjectID)

	return task, nil
}

func (ts *TaskService) GetAllTasks() ([]*models.Task, error) {
	findOptions := options.Find()
	var tasks  []*models.Task

	curr, err := ts.taskCollection.Find(context.Background(), bson.D{{}}, findOptions)
	if err != nil {
		return nil, err
	} 

	defer curr.Close(context.Background())

	for curr.Next(context.Background()) {
		var task models.Task
		err := curr.Decode(&task)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, &task)
	}

	if err := curr.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (ts *TaskService) GetTask(id string) (*models.Task, error){
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}
	var task models.Task

	err = ts.taskCollection.FindOne(context.Background(), filter).Decode(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (ts *TaskService) UpdateTask(id string, updatedTask models.Task) (error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	updatedFields := bson.M{}

	if updatedTask.Title != "" {
		updatedFields["title"] = updatedTask.Title
	}

	if updatedTask.Description != "" {
		updatedFields["description"] = updatedTask.Description
	}

	if !updatedTask.DueDate.IsZero() {
		updatedFields["due_date"] = updatedTask.DueDate
	}

	if updatedTask.Status != "" {
		updatedFields["status"] = updatedTask.Status
	}

	update := bson.M{
		"$set" : updatedFields,
	}

	_, err = ts.taskCollection.UpdateOne(context.Background(),filter, update)

	return err
}

func (ts *TaskService) DeleteTask(id string) (error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objectId}

	_, err = ts.taskCollection.DeleteOne(context.Background(), filter)

	return err
}