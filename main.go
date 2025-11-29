package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/SkyPromp/goLearning/controllers"

	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
	_ "github.com/SkyPromp/goLearning/docs"
)

var router *gin.Engine

func main(){
	router = gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/", controllers.RerouteToSwagger)

	router.GET("/todos", controllers.GetAll)
	router.GET("/todos/:id", controllers.GetById)
	router.POST("/todos", controllers.AddTodo)
	err := router.Run("localhost:9090")

	if err == nil{
		fmt.Println("Could not connect to localhost:9090")
	}
}

