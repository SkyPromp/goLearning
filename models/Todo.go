package models

type Todo struct {
	Id int `json:"id"`
	Task string `json:"title"`
	Completed bool `json:"completed"`
}
