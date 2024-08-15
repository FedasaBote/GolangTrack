package controllers

import (
	"clean_architecture/domain/dtos"
	"clean_architecture/domain/interface/usecase"
	"clean_architecture/domain/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


type TaskController struct {
	TaskUseCase usecase.TaskUseCase
}

func NewTaskController(taskUseCase usecase.TaskUseCase) TaskController {
	return TaskController{
		TaskUseCase: taskUseCase,
	}
}


func (t *TaskController) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := t.TaskUseCase.Create(c, task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (t *TaskController) GetAllTasks(c *gin.Context) {
	result, err := t.TaskUseCase.GetAllTasks(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (t *TaskController) GetTaskById(c *gin.Context) {
	id := c.Param("id")
	result, err := t.TaskUseCase.GetTaskById(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (t *TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task dtos.UpdateTaskDTO
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := t.TaskUseCase.UpdateTask(c, id, task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)

}

func (t *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := t.TaskUseCase.DeleteTask(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}