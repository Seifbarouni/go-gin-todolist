package models

type Todo struct {
	Todostr string
}
type todolist []*Todo

var todos todolist

func AddTodo(item *Todo) {
	todos = append(todos, item)
}

func GetTodos() *todolist {
	return &todos
}