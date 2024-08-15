package infrastructure

import (
	"clean_architecture/domain/dtos"
	"clean_architecture/domain/models"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type TaskRepository struct {
	database *mongo.Database
	collection string
}

func NewTaskRepository(database *mongo.Database, collection string) *TaskRepository {
	return &TaskRepository{database: database, collection: collection}
}


// Create a new task
func (t *TaskRepository) Create(c context.Context, task models.Task) (models.Task, error) {
	if task.DueDate.Before(time.Now()) {
		return models.Task{}, errors.New("due date must be in the future")
	}

	_, err := t.database.Collection(t.collection).InsertOne(c, task)

	if err != nil {
		return models.Task{}, err
	}

	return task, nil

}

// Get all tasks
func (t *TaskRepository) GetAllTasks(c context.Context) ([]models.Task, error) {
	var tasks []models.Task

	cursor, err := t.database.Collection(t.collection).Find(c, bson.M{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(c)

    err = cursor.All(c, &tasks)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// Get a task by id
func (t *TaskRepository) GetTaskById(c context.Context, id string) (models.Task, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return models.Task{}, err
	}

	var task models.Task

	err = t.database.Collection(t.collection).FindOne(c, bson.M{"_id": objectId}).Decode(&task)

	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (t *TaskRepository) UpdateTask(c context.Context, id string, task dtos.UpdateTaskDTO) (models.Task, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return models.Task{}, err
	}

	update := bson.M{}

	if task.Title != nil {
		update["title"] = *task.Title
	}

	if task.Description != nil {
		update["description"] = *task.Description
	}

	if task.DueDate != nil {
		dueDate, err := time.Parse(time.RFC3339, *task.DueDate)

		if err != nil {
			return models.Task{}, err
		}

		update["due_date"] = dueDate
	}

	if task.Status != nil {
		update["status"] = *task.Status
	}

	_, err = t.database.Collection(t.collection).UpdateOne(c, bson.M{"_id": objectId}, bson.M{"$set": update})

	if err != nil {
		return models.Task{}, err
	}

	return t.GetTaskById(c, id)
}

//A. We can use pointers to set optional fields in a struct in Go.
//for example, in the UpdateTaskDTO struct, we have used pointers to set optional fields.
// why it wokrs ?
// because a pointer to a type in Go can be nil which means it is not set.
// In the UpdateTaskDTO struct, we have used pointers to string to make the fields optional.
// When we get the value of the pointer, if it is nil, it means the field is not set.
// If it is not nil, it means the field is set.
// This way, we can set optional fields in a struct in Go.
// Let's see the UpdateTaskDTO struct:
// package dtos
// 
// // update task dto with optional fields
// type UpdateTaskDTO struct {
// 	Title *string `json:"title"`
// 	Description *string `json:"description"`
// 	DueDate *string `json:"due_date"`
// 	Status *string `json:"status"`
// }
//


func (t *TaskRepository) DeleteTask(c context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_, err = t.database.Collection(t.collection).DeleteOne(c, bson.M{"_id": objectId})

	if err != nil {
		return err
	}

	return nil
}