package services

import (
	"time"
	"errors"
	"github.com/SkyPromp/goLearning/models"
	"github.com/SkyPromp/goLearning/data"
)

func GetAll() ([]models.Todo){
	return data.Todos
}

func isIdEqual(value models.Todo, id int) (*models.Todo){
	time.Sleep(150 * time.Millisecond)

	if(value.Id == id){
		return &value
	}

	return nil
}

func GetById(id int) (*models.Todo, error){
	for _, value := range data.Todos {
		if result := isIdEqual(value, id); result != nil{
			return &value, nil
		}
	}

	return nil, errors.New("item not found")
}

func isIdEqualGoroutine(value models.Todo, id int, ch chan <- *models.Todo){
	time.Sleep(150 * time.Millisecond)

	if(value.Id == id){
		ch <- &value
	} else{
		ch <- nil
	}
}

func GetByIdGoroutine(id int) (*models.Todo, error){
	ch := make(chan *models.Todo)

	for _, value := range data.Todos {
		go isIdEqualGoroutine(value, id, ch)
	}

	for range data.Todos{
		if result :=<- ch; result != nil{
			return result, nil
		}
	}

	return nil, errors.New("item not found")
}

func AddTodo(value models.Todo) (models.Todo){
	max_id := -1

	for _, todo := range data.Todos{
		if todo.Id > max_id{
			max_id = todo.Id
		}
	}

	value.Id = max_id + 1
	data.Todos = append(data.Todos, value)

	return value
}
