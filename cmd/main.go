package main

import (
	"context"
	"fmt"
	"github.com/3XBAT/todo-app_by_yourself/configs"
	authgrpc "github.com/3XBAT/todo-app_by_yourself/pkg/clients/auth/grpc"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/3XBAT/todo-app_by_yourself"
	handler "github.com/3XBAT/todo-app_by_yourself/pkg/handlers"
	"github.com/3XBAT/todo-app_by_yourself/pkg/repository"
	"github.com/3XBAT/todo-app_by_yourself/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {

	cfg := configs.MustLoad()
	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	authClient, err := authgrpc.NewClient(
		context.Background(),
		log,
		cfg.Auth.Address,
		cfg.Auth.Timeout,
		cfg.Auth.RetriesCount,
	)

	db, err := repository.NewPostgresDB(repository.Config{
		Port:     "5432",
		Host:     "localhost",
		Username: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
		Password: "qwerty",
	})

	if err != nil {
		log.Error(fmt.Sprintf("failed to initialized db: %s", err.Error()))
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handler.NewHandler(services, authClient)

	srv := new(todo.Server)

	go func() {
		if err := srv.Run("8080", handler.InitRoutes()); err != nil {
			log.Error("error occured while runing the server %s","", err.Error())
		}
	}()

	logrus.Info("server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err := srv.ShutDown(context.Background()); err != nil {
		logrus.Errorf("error occured while shutting down server: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured while closing db: %s", err.Error())
	}
}
