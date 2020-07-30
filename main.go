package main

import (
	controllers "github.com/MgHtinLynn/golang-rest-api/controllers"
	"github.com/MgHtinLynn/golang-rest-api/models"
	"github.com/gin-gonic/gin"

)

func main() {

	router := gin.Default()

	db := models.SetupModels()

	// Provide db variable to controllers
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	v1 := router.Group("/api/v1/user")
	{
		v1.POST("/", controllers.CreateUser)
		v1.GET("/", controllers.FindUsers)
		v1.GET("/:id", controllers.FindUser)
		v1.PUT("/:id", controllers.UpdateUser)
		v1.DELETE("/:id", controllers.DeleteUser)

	}

	router.Run()

}
