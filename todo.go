package todo

import "errors"

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UsersList struct {
	Id      int
	User_id int
	List_id int
}

type TodoItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}

type ListItems struct {
	Id      int
	List_id int
	Item_id int
}

type UpdateListInput struct { //эта структура создаётся для того чтобы пользователь мог конкретно указать что он хочет обновить в своём списке
	Title       *string `json:"title"`
	Description *string `json:"description"`
	
}

func (i UpdateListInput) Validate() error { //мы не хотим запускать уровень репозитроия если у нас все значения для обновления - nil, для этого сделана эта функция
	if i.Title == nil && i.Description == nil {
		return errors.New("update sstructure has no values")
	}

	return nil
}

type UpdateItemInput struct { //эта структура создаётся для того чтобы пользователь мог конкретно указать что он хочет обновить в своём списке
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool    `json:"done" db:"done"`
}

func (i UpdateItemInput) Validate() error { //мы не хотим запускать уровень репозитроия если у нас все значения для обновления - nil, для этого сделана эта функция
	if i.Title == nil && i.Description == nil && i.Done == nil {
		return errors.New("update sstructure has no values")
	}

	return nil
}