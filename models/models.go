package models

type Todo struct {
	ID      int
	Todostr string
}
type todolist []*Todo

var todos todolist

func AddTodo(item *Todo) {
	todos = append(todos, item)
	update()
}

func GetTodos() *todolist {
	return &todos
}

func update() {
	for i, todo := range todos {
		todo.ID = i + 1
	}
}

func DeleteTodo(id int) {
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
		}
	}
	update()
}

func UpdateTodo(id int, todostr string) {
	for _, todo := range todos {
		if todo.ID == id {
			if todostr == "" {
				break
			} else {
				todo.Todostr = todostr
			}
		}
	}
	update()
}