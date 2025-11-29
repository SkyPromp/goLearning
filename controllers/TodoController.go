package controllers

import (
	"fmt"
	"time"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/SkyPromp/goLearning/models"
	"github.com/SkyPromp/goLearning/services"
)

func GetAll(context *gin.Context){
	start := time.Now()
	value := services.GetAll()
	duration := int64(time.Since(start))

	debugValue := models.Debug{Duration: duration, Data: value}

	context.IndentedJSON(http.StatusOK, debugValue)
}

func GetById(context *gin.Context){
	id, err := strconv.Atoi(context.Param("id"))
	usesGoRoutine, err2 := strconv.ParseBool(context.DefaultQuery("goroutine", "false"))

	if(err2 != nil){
		//usesGoRoutine = false
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Could not convert data to boolean"})
		return
	}

	if (err != nil) {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Id could not be converted to type: int"})
		return
	}

	start := time.Now()

	var value *models.Todo

	if(usesGoRoutine){
		value, err = services.GetByIdGoroutine(id)
	} else{
		value, err = services.GetById(id)
	}

	duration := int64(time.Since(start))

	if err != nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item not found"})
		return
	}

	debugValue := models.Debug{Duration: duration, Data: *value}

	context.IndentedJSON(http.StatusOK, debugValue)
}

func AddTodo(context *gin.Context){
	var value models.Todo

	if err := context.BindJSON(&value); err != nil {
		fmt.Println("Error has been found")
		return;
	}

	start := time.Now()
	services.AddTodo(value)
	duration := int64(time.Since(start))

	debugValue := models.Debug{Duration: duration, Data: value}

	context.IndentedJSON(http.StatusCreated, debugValue)
}

