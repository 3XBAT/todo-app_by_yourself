package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	Port     string
	Username string
	Host     string
	DBName   string
	Password string
	SSLMode  string
}

// возможно потому что создание бд никак не связано с взаимодейсвтием уровней, поэтому внедрять зависимость бессмысленно
func NewPostgresDB(cfg Config) (*sqlx.DB, error) { // создаём экземпляр прямо в методе. Почему это не противоречит чистой ахитектуре?????
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)) //

	if err != nil {
		return nil, err
	}

	err = db.Ping() // проверяем можем ли мы подколючится к нашей бд, если нет то кидаем ошибку
	if err != nil {
		return nil, err
	}
	fmt.Println("все ахуенно")
	return db, nil //если все ок возвращаем экземпляр нашей бд

}
