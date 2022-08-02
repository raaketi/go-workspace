package main

import (
	"io"
	"net/http"
	"os"

	"rajasureshaditya/go-workspace/Userapiservice/controllers"
	"rajasureshaditya/go-workspace/Userapiservice/models"

	"github.com/gin-gonic/gin"
)

var (
	modelinterface  models.Userserviceinterface      = models.Newmodel()
	usercontrollers controllers.Ecommercecontrollers = controllers.NewcntrlApp(modelinterface)
)

func main() {
	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	router := gin.Default()
	router.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "I am Healthy"})
	})
	router.GET("/GetUsers", usercontrollers.GetUserscntrl)
	router.POST("/Createuser", usercontrollers.Createnewuser)
	router.Run()
}
