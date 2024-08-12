package main

import (
	"context"
	"task_manager/data"
	"task_manager/router"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	db, err := getDbConnection()
	if err != nil {
		panic(err)
	}

	taskCollection := db.Collection("tasks")
	taskService := data.NewTaskService(taskCollection)
	r := router.SetupRouter(*taskService)
	r.Run(":8080")


}

func getDbConnection()(*mongo.Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	err = client.Connect(context.Background())

	if err != nil {
		panic(err)
	}

	db := client.Database("task_manager")
	return db, nil
	
}