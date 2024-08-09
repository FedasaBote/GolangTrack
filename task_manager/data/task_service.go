package data

import (
	"errors"
	"task_manager/models"
	"time"
)


type TaskService struct {
	tasks []models.Task
}

func NewTaskService() *TaskService {
	return &TaskService{
		tasks: make([]models.Task, 0),
	}
}

func (ts *TaskService) AddTask(task models.Task) (err error) {
	if task.DueDate.Before(time.Now()) {
		return errors.New("due date can't be in the past")
	}

	for _,otherTask := range ts.tasks {
		if otherTask.ID == task.ID {
			return errors.New("task with the given Id already exists")
		}
	}

	ts.tasks = append(ts.tasks, task)
	return nil
}

func (ts *TaskService) GetAllTasks() []models.Task {
	return ts.tasks
}

func (ts *TaskService) GetTask(id string) (task models.Task, err error){
	
	for _, task := range ts.tasks {
		if task.ID == id {
			return task, nil
		}
	}

	return models.Task{}, errors.New("task not found");

}

func (ts *TaskService) UpdateTask(id string, updatedTask models.Task) (err error) {

	for i, task := range ts.tasks {
		if task.ID == id {
			if updatedTask.Description != ""  {
				task.Description = updatedTask.Description
			}

			if updatedTask.DueDate.After(time.Now()){
				task.DueDate = updatedTask.DueDate
			}
			
			if updatedTask.Status != "" {
				task.Status = updatedTask.Status
			}
			ts.tasks[i] = task
		return nil
		}
	}

	return errors.New("task not found")
}

func (ts *TaskService) DeleteTask(id string) (err error) {

	for i, task := range ts.tasks {
		if task.ID == id {
			ts.tasks = append(ts.tasks[:i], ts.tasks[i + 1:]...)
			return nil
		}
	}

	return errors.New("task not found")
}