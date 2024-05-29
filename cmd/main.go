package main

import (
	"fmt"
	"log"
	"os"

	//"net/http"
	"github.com/3XBAT/todo-app_by_yourself"
	handler "github.com/3XBAT/todo-app_by_yourself/pkg/handlers"
	"github.com/3XBAT/todo-app_by_yourself/pkg/repository"
	"github.com/3XBAT/todo-app_by_yourself/pkg/service"
	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config : %s", err.Error())
	} 
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := gotenv.Load(); err != nil {
		log.Fatalf("failed loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{ 
		Port:     viper.GetString("db.port"),
		Host:     viper.GetString("db.host"),
		Username: viper.GetString("db.name"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	}) 

	if err != nil {
		log.Fatalf(fmt.Sprintf("failed to initialized db: %s", err.Error()))
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	
	srv := new(todo.Server)

	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil { 
		log.Fatalf("error occured while runing the server %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("CONFIGS") 
	viper.SetConfigFile("config.yaml")
	return viper.ReadInConfig()
}

// важное замечание, если перменная и папка названны одинаково, то это очень плохо, т.к. когда в коде встречается запись "имя_папки_и_переменной". то происходит конфлик имён и наш компилятор хз что делать
