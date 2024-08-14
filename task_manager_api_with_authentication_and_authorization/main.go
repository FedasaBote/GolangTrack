package main

import (
	"context"
	"enhanced_task_manager/data"
	"enhanced_task_manager/router"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	db, err := getDBConnection()

	if err != nil {
		log.Fatal(err)
	}

	taskCollection := db.Collection("tasks")
	userCollection := db.Collection("users")

	taskService := data.NewTaskService(taskCollection)
	userService := data.NewUserService(userCollection)
	r := router.SetupRouter(*taskService, *userService)
	r.Run(":8080")
}

func connectToMongoDB(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB")
    return client, nil
}

func getDBConnection() (*mongo.Database, error) {
	client, err := connectToMongoDB("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	db := client.Database("taskdb")

	return db, nil
}