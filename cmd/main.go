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
	} //почему сначала инициализации конфига а потом все остальное?-чтобы мы могли юзать viper.GetString() при создании экземпляра БД

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := gotenv.Load(); err != nil {
		log.Fatalf("failed loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{ //для того чтобы компилятор увидел функцию NewPostgresDB нужно четко указать это через точку
		Port:     viper.GetString("db.port"),
		Host:     viper.GetString("db.host"),
		Username: viper.GetString("db.name"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	}) //наша бд инициализируется значениями из конфига, поэтому мы сначала инициализируем без

	if err != nil {
		log.Fatalf(fmt.Sprintf("failed to initialized db: %s", err.Error()))
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	//handler := new(handler.Handler) //сам по себе handler.Handler это тип, ни конструктор, ни функция, поэтому мы юзаем функцию new чтобы создать объект этого типа
	srv := new(todo.Server) // тоже самое

	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil { //handler.InitRoutes() подходт,т.к. эта фун-ция возвращает gin.Engine объект, который реализует интерфейс хендлера из пакета http
		log.Fatalf("error occured while runing the server %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("CONFIGS") //можно и без неё, нахуй она не нужна впринципе
	viper.SetConfigFile("config.yaml")
	return viper.ReadInConfig()
}

// важное замечание, если перменная и папка названны одинаково, то это очень плохо, т.к. когда в коде встречается запись "имя_папки_и_переменной". то происходит конфлик имён и наш компилятор хз что делать
