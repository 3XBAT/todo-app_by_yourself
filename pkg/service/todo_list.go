package service

import (
	"github.com/3XBAT/todo-app_by_yourself"
	"github.com/3XBAT/todo-app_by_yourself/pkg/repository"
)

type TodoListService struct {//почему мы создаём свой сервис для работы с разными сущностями? для того чтобы разделить отвественность
	repo repository.TodoList// потому, что т.о. мы разделяем между ними отвественность
	//именно эта структура реализовывает все методы TodoList интерфейса, никакая другая и если что
	// мы можем сделать одну структуру которая будер реализовывать все методы всех интерфейсов
	// но это само по себе плохая идея(SRP),потому что изменения в одном методе могут повлечь за собой изменения в другом
	// 
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo:repo}
}

func (s *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	//задача этого метода переслать все данные в репозиторий и вызвать там соотвествующий метод
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]todo.TodoList, error){
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId, listId int) (todo.TodoList, error){
	return s.repo.GetById(userId, listId)
}

func (s *TodoListService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *TodoListService) Update(userId, listId int, input todo.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, listId, input)
}