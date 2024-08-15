package usecase

import (
	"clean_architecture/domain/interface/repository"
	"clean_architecture/domain/interface/usecase"
	"clean_architecture/domain/models"
	"context"
	"time"
)

type userUsecase struct {
	userRepository repository.UserRepository
	contextTimeout time.Duration
}

func NewUserUseCase(userRepository repository.UserRepository, timeout time.Duration) usecase.UserUseCase {
	return &userUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (uu *userUsecase) Create(c context.Context, user models.User) (models.User,error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	return uu.userRepository.Create(ctx, user)
}

func (uu *userUsecase) Login(c context.Context, user models.User)(string, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	return uu.userRepository.Login(ctx, user)
}

func (uu *userUsecase) Promote(c context.Context, id string) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	return uu.userRepository.Promote(ctx, id)
}