package main

import (
	"log"
	//"net/http"

	"github.com/3XBAT/todo-app_by_yourself"
	handler "github.com/3XBAT/todo-app_by_yourself/pkg/handlers"
	"github.com/3XBAT/todo-app_by_yourself/pkg/repository"
	"github.com/3XBAT/todo-app_by_yourself/pkg/service"
	//"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil{
		log.Fatalf("error initializing config : %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	//handler := new(handler.Handler) //сам по себе handler.Handler это тип, ни конструктор, ни функция, поэтому мы юзаем функцию new чтобы создать объект этого типа
	srv := new(todo.Server) // тоже самое

	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil { //handler.InitRoutes() подходт,т.к. эта фун-ция возвращает gin.Engine объект, который реализует интерфейс хендлера из пакета http
		log.Fatalf("error occured while runing the server %s", err.Error())
	}

}

func initConfig() error{
	viper.AddConfigPath("configs")
	viper.SetConfigFile("config")
	return viper.ReadInConfig()
}

// важное замечание, если перменная и папка названны одинаково, то это очень плохо, т.к. когда в коде встречается запись "имя_папки_и_переменной". то происходит конфлик имён и наш компилятор хз что делать
