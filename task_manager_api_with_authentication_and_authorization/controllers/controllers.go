package controllers

import (
	"enhanced_task_manager/data"
	"enhanced_task_manager/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	taskService *data.TaskService
	userService *data.UserService
}

func NewController(taskService data.TaskService, userService data.UserService) *Controller {
	return &Controller{
		taskService: &taskService,
		userService: &userService,
	}
}

func (ctr *Controller) AddTask(c *gin.Context) {
	var task models.Task

	if err := c.BindJSON(&task); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid Request."})
		return
	}

	task, err := ctr.taskService.AddTask(task);
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, task)
}

func (ctr *Controller) GetAllTasks(c *gin.Context) {
	task, err := ctr.taskService.GetAllTasks()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, task)
}

func (ctr *Controller) GetTask(c *gin.Context) {
	id := c.Param("id")
	task, err := ctr.taskService.GetTask(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found!"})
		return
	}

	c.IndentedJSON(http.StatusOK, task)
}


func (ctr *Controller) UpdateTask(c *gin.Context) {
	var updatedTask models.Task
	id := c.Param("id")

	if err := c.BindJSON(&updatedTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	err := ctr.taskService.UpdateTask(id, updatedTask);

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task Not found"})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message" : "Task updated successfully!"})
}

func (ctr *Controller) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	err := ctr.taskService.DeleteTask(id);
	
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task Not Found"})
		return
	}

	c.IndentedJSON(http.StatusNoContent, gin.H{"message": "Task Deleted Succesfully!"})
}


func (ctr *Controller) RegisterUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}

	err := ctr.userService.RegisterUser(user)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "User registerd successfully!"})
}

func (ctr *Controller) Login(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}

	jwtToken, loginError := ctr.userService.Login(user);
	if loginError != nil {
		c.IndentedJSON(http.StatusUnauthorized, loginError.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "User logged in successfully!", "token": jwtToken})
}

func (ctr *Controller) Promote(c *gin.Context) {
	id := c.Param("id")

	err := ctr.userService.Promote(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found!"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "User promoted successfully!"})
}