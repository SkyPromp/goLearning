package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/SkyPromp/goLearning/controllers"
)

func main(){
	router := gin.Default()
	router.GET("/", controllers.GetRoot)
	router.GET("/todos", controllers.GetAll)
	router.GET("/todos/:id", controllers.GetById)
	router.POST("/todos", controllers.AddTodo)
	err := router.Run("localhost:9090")

	if err == nil{
		fmt.Println("Could not connect to localhost:9090")
	}
}
