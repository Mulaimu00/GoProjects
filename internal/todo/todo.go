package todo

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

func (svc *Service) Remove(id int) {
	svc.todos = append(svc.todos[:id], svc.todos[id+1:]...)
}

func (svc *Service) GetAll() []string {
	return svc.todos
}
