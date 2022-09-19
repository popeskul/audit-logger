package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/popeskul/audit-logger/pkg/config"
	"log"
	"os"
	"time"

	"github.com/popeskul/audit-logger/pkg/handler"
	"github.com/popeskul/audit-logger/pkg/repository"
	"github.com/popeskul/audit-logger/pkg/server"
	"github.com/popeskul/audit-logger/pkg/service"
)

func main() {
	cfg, err := initConfig("config")
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := repository.NewMongoDB(ctx, repository.Config{
		Username: cfg.DB.DBUsername,
		Password: cfg.DB.DBPassword,
		Url:      fmt.Sprintf("mongodb://%s:%d", cfg.DB.Host, cfg.DB.Port),
	})
	if err != nil {
		log.Fatalln(err)
	}

	repo := repository.NewRepository(db)
	service := service.NewLogService(repo)
	handlers := handler.NewHandler(service)
	srv := server.NewServer(handlers, server.LoggingUnaryInterceptor)

	log.Fatalln(srv.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port)))
}

func initConfig(filename string) (*config.Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	cfg, err := config.New("configs", filename)
	if err != nil {
		return nil, err
	}

	cfg.DB.DBPassword = os.Getenv("DB_PASSWORD")
	cfg.DB.DBUsername = os.Getenv("DB_USERNAME")

	return cfg, nil
}
