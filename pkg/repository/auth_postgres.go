package repository

import (
	"fmt"

	"github.com/3XBAT/todo-app_by_yourself"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB //дальше некуда
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err 
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (todo.User, error) {
	var user todo.User
	query := fmt.Sprintf("SELECT id FROM %s where username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user,query,username, password)// почему мы тут не делаем провекру на ошибки -- вопрос открытый
	
	return user, err
}