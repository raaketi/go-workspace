package controllers

import (
	"rajasureshaditya/go-workspace/userservice/models"

	"github.com/gin-gonic/gin"
)

type UserControllerService struct {
	Userservice models.UserInterface
}

type UserControllers interface {
	Createuser(ctx *gin.Context) models.User
	GetAllusers() []models.User
}

func NewApp(userservice models.UserInterface) UserControllers {
	return &UserControllerService{
		Userservice: userservice,
	}
}

func (us *UserControllerService) GetAllusers() []models.User {
	return us.Userservice.GetAllusers()
}

func (us *UserControllerService) Createuser(ctx *gin.Context) models.User {
	var User models.User
	ctx.ShouldBindJSON(&User)
	us.Userservice.Createuser(User)
	return User
}
