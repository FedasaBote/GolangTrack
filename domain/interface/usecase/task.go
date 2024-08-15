package usecase

import (
	"clean_architecture/domain/dtos"
	"clean_architecture/domain/models"
	"context"
)



type TaskUseCase interface {
	Create(c context.Context,task models.Task) (models.Task,error)
	GetAllTasks(c context.Context)([]models.Task,error)
	GetTaskById(c context.Context,id string)(models.Task,error)
	UpdateTask(c context.Context,id string,task dtos.UpdateTaskDTO)(models.Task,error)
	DeleteTask(c context.Context,id string)error
}