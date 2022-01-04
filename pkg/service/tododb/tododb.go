package tododb

var (
	TodoListInstace = TodoList{}
)

type Todo struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsDone      bool   `json:"is_done"`
}

type TodoList struct {
	todos []Todo
}

func (t *TodoList) Add(todo Todo) Todo {
	id := len(t.todos) + 1
	todo.Id = id
	t.todos = append(t.todos, todo)
	return todo
}

func (t *TodoList) List() []Todo {
	return t.todos
}
