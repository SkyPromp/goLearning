package main

import (
	"fmt"

	"github.com/SkyPromp/goLearning/controllers"
	"github.com/gin-gonic/gin"

	_ "github.com/SkyPromp/goLearning/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var router *gin.Engine

func main(){
	router = gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/", controllers.RerouteToSwagger)

	router.GET("/todos", controllers.GetAll)
	router.GET("/todos/:id", controllers.GetById)
	router.POST("/todos", controllers.AddTodo)

	router.GET("/memory", controllers.GetMemoryManagementExample)
	router.GET("/alignment", controllers.GetByteAlignment)
	router.GET("/unsafe", controllers.GetUnsafeExample)

	err := router.Run("localhost:9090")

	if err == nil{
		fmt.Println("Could not connect to localhost:9090")
	}
}

