package infrastructure

import (
	"clean_architecture/domain/models"
	"clean_architecture/utils"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)



type UserRepository struct {
	database *mongo.Database
	collection string
}

func NewUserRepository(database *mongo.Database, collection string) *UserRepository {
	return &UserRepository{database: database, collection: collection}
}

// Create a new user
func (u *UserRepository) Create(c context.Context, user models.User) (models.User, error) {
	// hash the password
	hashedPassword, err := utils.GeneratePasswordHash(user.Password)
	if err != nil {
		return models.User{}, err
	}

	user.Password = hashedPassword

	// if its the first user, promote to admin
	count, err := u.database.Collection(u.collection).CountDocuments(c, bson.M{})
	if err != nil {
		return models.User{}, err
	}

	if count == 0 {
		user.Role = "admin"
	}else {
		user.Role = "user"
	}

	// check if the email is already in use
	var existingUser models.User
	err = u.database.Collection(u.collection).FindOne(c, bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		return models.User{}, errors.New("email already in use")
	}

	_, err = u.database.Collection(u.collection).InsertOne(c, user)

	if err != nil {
		return models.User{}, err
	}

	return user, nil

}


// Login a user
func (u *UserRepository) Login(c context.Context, user models.User) (string, error) {
	var existingUser models.User
	err := u.database.Collection(u.collection).FindOne(c, bson.M{"email": user.Email}).Decode(&existingUser)
	if err != nil {
		return "", errors.New("invalid email")
	}

	if !utils.ComparePasswordHash(user.Password, existingUser.Password) {
		return "", errors.New("invalid password")
	}
	jwtSecret := utils.GetJwtSecret("JWT_SECRET")
	jwtGenerator := utils.NewJWTService([]byte(jwtSecret))

	token, err := jwtGenerator.GenerateToken(existingUser.ID, existingUser.Email,existingUser.Role, 24*60*60)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *UserRepository) Promote(c context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_, err = u.database.Collection(u.collection).UpdateOne(c, bson.M{"_id": objectId}, bson.M{"$set": bson.M{"role": "admin"}})

	if err != nil {
		return err
	}

	return nil
}
