package router

import (
	"task_manager/controllers"
	"task_manager/data"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	taskService := *data.NewTaskService()
	taskController := *controllers.NewTaskController(taskService)

	r.GET("/tasks", taskController.GetAllTasks)
	r.GET("/tasks/:id", taskController.GetTask)
	r.PUT("/tasks/:id", taskController.UpdateTask)
	r.POST("tasks", taskController.AddTask)
	r.DELETE("/tasks/:id", taskController.DeleteTask)

	return r
}