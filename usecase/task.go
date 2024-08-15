package usecase

import (
	"clean_architecture/domain/dtos"
	"clean_architecture/domain/interface/repository"
	"clean_architecture/domain/interface/usecase"
	"clean_architecture/domain/models"
	"context"
	"time"
)

type taskUsecase struct {
	taskRepository repository.TaskRepository
	contextTimeout time.Duration
}

func NewTaskUseCase(taskRepository repository.TaskRepository, timeout time.Duration) usecase.TaskUseCase {
	return &taskUsecase{
		taskRepository: taskRepository,
		contextTimeout: timeout,
	}
}


func (tu *taskUsecase) Create(c context.Context, task models.Task) (models.Task,error){
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	return tu.taskRepository.Create(ctx, task)
}

func (tu *taskUsecase) GetAllTasks(c context.Context)([]models.Task, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	return tu.taskRepository.GetAllTasks(ctx)
}

func (tu *taskUsecase) GetTaskById(c context.Context, id string)(models.Task, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	return tu.taskRepository.GetTaskById(ctx, id)
}

func (tu *taskUsecase) UpdateTask(c context.Context, id string, task dtos.UpdateTaskDTO) (models.Task,error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	return tu.taskRepository.UpdateTask(ctx, id, task)
}

func (tu *taskUsecase) DeleteTask(c context.Context, id string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	return tu.taskRepository.DeleteTask(ctx, id)
}