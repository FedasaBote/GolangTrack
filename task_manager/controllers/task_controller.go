package controllers

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskService *data.TaskService
}

func NewTaskController(taskService data.TaskService) *TaskController {
	return &TaskController{
		taskService: &taskService,
	}
}

func (ts *TaskController) AddTask(c *gin.Context) {
	var task models.Task

	if err := c.BindJSON(&task); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid Request."})
		return
	}

	_, err := ts.taskService.AddTask(task)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Task Created"})
}

func (ts *TaskController) GetAllTasks(c *gin.Context) {
	tasks, err := ts.taskService.GetAllTasks()

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No Task Found!"})
		return
	}

	c.IndentedJSON(http.StatusOK, tasks)
}

func (ts *TaskController) GetTask(c *gin.Context) {
	id := c.Param("id")
	task, err := ts.taskService.GetTask(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found!"})
		return
	}

	c.IndentedJSON(http.StatusOK, task)
}


func (ts *TaskController) UpdateTask(c *gin.Context) {
	var updatedTask models.Task
	id := c.Param("id")

	if err := c.BindJSON(&updatedTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	err := ts.taskService.UpdateTask(id, updatedTask);

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task Not found"})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message" : "Task updated successfully!"})
}

func (ts *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	err := ts.taskService.DeleteTask(id);
	
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task Not Found"})
		return
	}

	c.IndentedJSON(http.StatusNoContent, gin.H{"message": "Task Deleted Succesfully!"})
}