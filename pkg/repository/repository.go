package repository

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository interface {
	Authorization
	TodoItem
	TodoList
}

func NewRepository() *Repository {
	return &Repository{}
}
