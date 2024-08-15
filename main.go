package main

import (
	"clean_architecture/api"
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func main() {
	db, err := getDBConnection()

	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	api.Router( db, r)

	r.Run(":8080")

}

func getDBConnection()(*mongo.Database,error){
	client:= options.Client().ApplyURI("mongodb://localhost:27017")
	
	mongoConn, err := mongo.Connect(context.Background(), client)
	
	if err != nil {
		return nil, err
	}

	err = mongoConn.Ping(context.Background(), nil)

	if err != nil {
		return nil, err
	}

	db := mongoConn.Database("taskdb")

	return db, nil

}