package data

import (
	"context"
	"enhanced_task_manager/models"
	"errors"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userCollection *mongo.Collection
}


func NewUserService(userCollection *mongo.Collection) *UserService{
	return &UserService{
		userCollection: userCollection,
	}
}


func (us *UserService) RegisterUser(user models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("internal server error: failed to hash password")
	}

	count, err := us.userCollection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		return errors.New("internal server error: failed to count documents")
	}

	if count == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}

	filter := bson.M{"email": user.Email}
	err = us.userCollection.FindOne(context.Background(), filter).Err()

	if err == mongo.ErrNoDocuments {
		user.Password = string(hashedPassword)
		_, insertErr := us.userCollection.InsertOne(context.Background(), user)
		if insertErr != nil {
			return errors.New("internal server error: failed to insert user")
		}
		return nil
	}

	if err == nil {
		return fmt.Errorf("user already exists")
	}

	return errors.New("internal server error: failed to check if user exists")
}


func (us *UserService) Login(user models.User)(string, error) {
	filter := bson.M{"email": user.Email}
	var existingUser models.User

	err := us.userCollection.FindOne(context.Background(), filter).Decode(&existingUser)
	if err != nil {
		return "", err
	}

	var jwtSecret = []byte(getJwtSecret("JWT_SECRET"))

	if bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)) != nil {
		return "", errors.New("invalid email or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id": existingUser.ID,
			"email": existingUser.Email,
			"role": existingUser.Role,
	})

	jwtToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", errors.New("inernal serever error")
	}

	return jwtToken, nil
}

func (us *UserService) Promote(id string) (error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$set": bson.M{
			"role": "admin",
		},
	}

	_, err = us.userCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func getJwtSecret(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		return ""
	}

	return os.Getenv(key)
}
