package repository

import (
	"clean_architecture/domain/models"
	"context"
)


type UserRepository interface {
	Create(c context.Context, user models.User) (models.User,error)
	Login(c context.Context, user models.User) (string, error)
	Promote(c context.Context, id string) error

}