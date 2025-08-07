package todo

import "errors"

type Service struct {
	todos []string
}

func NewService() *Service {
	return &Service{
		todos: make([]string, 0),
	}
}

func (svc *Service) Add(todo string) {
	svc.todos = append(svc.todos, todo)
}

func (svc *Service) Remove(id int) error {
	if id < 0 || id >= len(svc.todos) {
		return errors.New("invalid id: out of range")
	}
	svc.todos = append(svc.todos[:id], svc.todos[id+1:]...)
	return nil
}

func (svc *Service) GetAll() []string {
	return svc.todos
}

func (svc *Service) Len() int {
	return len(svc.todos)
}
