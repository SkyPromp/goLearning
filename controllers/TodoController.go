package controllers

import (
	"fmt"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/SkyPromp/goLearning/models"
	"github.com/SkyPromp/goLearning/services"
)

func GetAll(context *gin.Context){
	context.IndentedJSON(http.StatusOK, services.GetAll())
}

func GetByIdHandler(context *gin.Context){
	id, err := strconv.Atoi(context.Param("id"))

	if (err != nil) {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Id could not be converted to type: int"})
		return
	}

	value, err := services.GetById(id)

	if err != nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, value)
}

func AddTodo(context *gin.Context){
	var value models.Todo

	if err := context.BindJSON(&value); err != nil {
		fmt.Println("Error has been found")
		return;
	}

	services.AddTodo(value)

	context.IndentedJSON(http.StatusCreated, value)
}

