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

func NewPostgresDB(cfg Config) (*sqlx.DB, error) { 
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)) //

	if err != nil {
		return nil, err
	}

	err = db.Ping() 
	if err != nil {
		return nil, err
	}
	fmt.Println("все ахуенно")
	return db, nil 

}
