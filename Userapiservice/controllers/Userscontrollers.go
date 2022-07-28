package controllers

import (
	"log"
	"net/http"

	"rajasureshaditya/go-workspace/Userapiservice/models"

	"github.com/gin-gonic/gin"
)

type Ecommercecontrollers interface {
	GetUserscntrl(c *gin.Context)
	Createnewuser(ctx *gin.Context)
}

type ecommercecntrlStruct struct {
	ModelInterface models.Userserviceinterface
}

func NewcntrlApp(localclientinterface models.Userserviceinterface) Ecommercecontrollers {
	return &ecommercecntrlStruct{
		ModelInterface: localclientinterface,
	}
}

func (modelintr *ecommercecntrlStruct) GetUserscntrl(c *gin.Context) {
	log.Println(modelintr.ModelInterface.GetUsers())
	c.JSON(http.StatusOK, gin.H{"Users": modelintr.ModelInterface.GetUsers()})
	// return c
}

// func (modelintr *ecommercecntrlStruct) GetUser() gin.HandlerFunc {
// 	log.Println(modelintr.ModelInterface.GetUsers())
// 	return func(ctx *gin.Context) {
// 		c, cancel := ctx.WithTimeout(ctx.Background(), 100*time.Second)
// 		c.JSON(http.StatusOK, gin.H{"Users": modelintr.ModelInterface.GetUsers()})
// 		defer cancel()
// 	}
// }

func (modelintr *ecommercecntrlStruct) Createnewuser(ctx *gin.Context) {
	var User *models.User
	ctx.ShouldBindJSON(&User)
	Users := *modelintr.ModelInterface.Createuser(User)
	ctx.JSON(http.StatusOK, gin.H{"Users": Users})
}
