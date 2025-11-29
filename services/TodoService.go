package services

import (
	"errors"
	"github.com/SkyPromp/goLearning/models"
	"github.com/SkyPromp/goLearning/data"
)

func GetAll() ([]models.Todo){
	return data.Todos
}

func GetById(id int) (*models.Todo, error){
	for _, value := range data.Todos {
		if value.Id == id{
			return &value, nil
		}
	}

	return nil, errors.New("item not found")
}

func AddTodo(value models.Todo){
	data.Todos = append(data.Todos, value)
}
