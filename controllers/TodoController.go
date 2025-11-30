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

// GetAll returns all todos with execution duration
// @Summary Get all todos
// @Description Returns all todos with debug info
// @Success 200 {object} models.Debug
// @Router /todos [get]
func GetAll(context *gin.Context){
	start := time.Now()
	value := services.GetAll()
	duration := time.Since(start).Nanoseconds()

	debugValue := models.Debug{Duration: duration, Data: value}

	context.IndentedJSON(http.StatusOK, debugValue)
}

// GetById returns a todo by ID, with optional goroutine query param
// @Summary Get todo by ID
// @Description Returns a single todo item by its ID. Optional query parameter 'goroutine' runs in a goroutine if true.
// @Tags Gouroutine example
// @Param id path int true "Todo ID"
// @Param goroutine query bool false "Use goroutine"
// @Success 200 {object} models.Debug
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /todos/{id} [get]
func GetById(context *gin.Context){
	id, err := strconv.Atoi(context.Param("id"))

	if (err != nil) {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Id could not be converted to type: int"})

		return
	}

	usesGoRoutine, err := strconv.ParseBool(context.DefaultQuery("goroutine", "false"))

	if(err != nil){
		usesGoRoutine = false
	}

	start := time.Now()

	var value *models.Todo

	if(usesGoRoutine){
		value, err = services.GetByIdGoroutine(id)
	} else{
		value, err = services.GetById(id)
	}

	duration := time.Since(start).Nanoseconds()

	if err != nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item not found"})

		return
	}

	debugValue := models.Debug{Duration: duration, Data: *value}

	context.IndentedJSON(http.StatusOK, debugValue)
}

// AddTodo adds a new todo item
// @Summary Add a new todo
// @Description Creates a new todo item. Request body must include the todo details.
// @Accept json
// @Produce json
// @Param todo body models.Todo true "Todo object"
// @Success 201 {object} models.Debug
// @Failure 400 {object} map[string]string
// @Router /todos [post]
func AddTodo(context *gin.Context){
	var value models.Todo

	if err := context.BindJSON(&value); err != nil {
		fmt.Println("Error has been found")
		return;
	}

	start := time.Now()
	newTodo := services.AddTodo(value)
	duration := time.Since(start).Nanoseconds()

	debugValue := models.Debug{Duration: duration, Data: newTodo}

	context.IndentedJSON(http.StatusCreated, debugValue)
}

