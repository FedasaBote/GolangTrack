package api

import (
	"clean_architecture/api/controllers"
	"clean_architecture/domain/models"
	"clean_architecture/infrastructure"
	"clean_architecture/usecase"
	"clean_architecture/utils"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)



func SetupRouter(timeout time.Duration, db *mongo.Database, gin *gin.Engine) {
	tr := infrastructure.NewTaskRepository(db, models.CollectionTask)
	ur := infrastructure.NewUserRepository(db, models.CollectionUser)
	userctr := &controllers.UserController{
		UserUseCase: usecase.NewUserUseCase(ur, timeout),
	}
	taskctr := &controllers.TaskController{
		TaskUseCase: usecase.NewTaskUseCase(tr, timeout),
	}
	
	publicRouter := gin.Group("")
	{
		publicRouter.POST("/register", userctr.CreateUser)
		publicRouter.POST("/login", userctr.Login)
	}

	userRouter := gin.Group("")
	userRouter.Use(utils.AuthMiddleWare())
	{
		userRouter.GET("/tasks", taskctr.GetAllTasks)
		userRouter.GET("tasks/:id", taskctr.GetTaskById)
	}

	adminRouter := gin.Group("")
	adminRouter.Use(utils.AuthMiddleWare(), utils.RoleMiddleware())
	{
		adminRouter.POST("/tasks", taskctr.CreateTask)
		adminRouter.PUT("/tasks/:id", taskctr.UpdateTask)
		adminRouter.DELETE("/tasks/:id", taskctr.DeleteTask)
		adminRouter.PUT("/users/promote/:id", userctr.Promote)
	}
}