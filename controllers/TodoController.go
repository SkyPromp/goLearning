package controllers

import (
	"fmt"
	"errors"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
)

type todo struct {
	Id int `json:"id"`
	Task string `json:"title"`
	Completed bool `json:"completed"`
}

var todos = []todo{
	{Id: 1, Task: "Do something", Completed:false},
	{Id: 2, Task: "Do nothing", Completed:false},
	{Id: 3, Task: "Do something else", Completed:false},
	{Id: 4, Task: "Do whatever", Completed:false},
}

func GetAll(context *gin.Context){
	context.IndentedJSON(http.StatusOK, todos)
}

func GetByIdHandler(context *gin.Context){
	id, err := strconv.Atoi(context.Param("id"))

	if (err != nil) {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Id could not be converted to type: int"})
		return
	}

	value, err := getById(id)

	if err != nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, value)
}

func getById(id int) (*todo, error){
	for _, value := range todos {
		if value.Id == id{
			return &value, nil
		}
	}

	return nil, errors.New("item not found")
}

func AddTodo(context *gin.Context){
	var value todo

	if err := context.BindJSON(&value); err != nil {
		fmt.Println("Error has been found")
		return;
	}

	todos = append(todos, value)

	context.IndentedJSON(http.StatusCreated, value)
}

