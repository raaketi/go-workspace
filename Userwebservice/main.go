package main

import (
	"log"
	"os"
	"rajasureshaditya/go-workspace/userservice/controllers"
	"rajasureshaditya/go-workspace/userservice/models"

	"github.com/gin-gonic/gin"
)

var (
	userimpletation models.UserInterface        = models.New()
	Usercontrollers controllers.UserControllers = controllers.NewApp(userimpletation)
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	routes := gin.New()

	routes.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, Usercontrollers.GetAllusers())
	})
	routes.POST("/createuser", func(ctx *gin.Context) {
		ctx.JSON(200, Usercontrollers.Createuser(ctx))
	})
	log.Printf("Listening on port %s", port)
	routes.Run(":" + port)

}
