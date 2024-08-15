package controllers

import (
	"clean_architecture/domain/interface/usecase"
	"clean_architecture/domain/models"
	"net/http"

	"github.com/gin-gonic/gin"
)



type UserController struct {
	UserUseCase usecase.UserUseCase
}

func NewUserController(u usecase.UserUseCase) UserController {
	return UserController{
		UserUseCase: u,
	}
}

func (u *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := u.UserUseCase.Create(c, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (u *UserController) Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := u.UserUseCase.Login(c, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (u *UserController) Promote(c *gin.Context) {
	id := c.Param("id")
	err := u.UserUseCase.Promote(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User promoted"})
}