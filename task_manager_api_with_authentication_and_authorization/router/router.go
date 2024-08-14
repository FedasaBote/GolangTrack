package router

import (
	"enhanced_task_manager/controllers"
	"enhanced_task_manager/data"
	"enhanced_task_manager/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(taskService data.TaskService, userService data.UserService) *gin.Engine {
	r := gin.Default()

	controller := *controllers.NewController(taskService, userService)

	// Public routes
	r.POST("/register", controller.RegisterUser)
	r.POST("/login", controller.Login)

	//Authenticated user routes
	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleWare())
	{
		authorized.GET("/tasks", controller.GetAllTasks)
		authorized.GET("/tasks/:id", controller.GetTask)
	}

	// Admin routes
	admin := r.Group("/")
	admin.Use(middleware.AuthMiddleWare(), middleware.RoleMiddleware())
	{
		admin.PUT("/tasks/:id", controller.UpdateTask)
		admin.POST("tasks", controller.AddTask)
		admin.DELETE("/tasks/:id", controller.DeleteTask)
		admin.PUT("/users/promote/:id", controller.Promote)
	}

	return r
}