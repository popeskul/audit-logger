package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/popeskul/audit-logger/pkg/client"
	"github.com/popeskul/audit-logger/pkg/config"
	"github.com/popeskul/audit-logger/pkg/domain"
	"github.com/popeskul/audit-logger/pkg/handler"
	"github.com/popeskul/audit-logger/pkg/repository"
	"github.com/popeskul/audit-logger/pkg/server"
	"github.com/popeskul/audit-logger/pkg/service"
	"github.com/popeskul/audit-logger/pkg/util"
)

var (
	loggerClient *client.Client
	cfg          *config.Config
)

func TestMain(m *testing.M) {
	if err := util.ChangeDir("../../"); err != nil {
		log.Fatal(err)
	}

	cfg1, err := initConfig("test.config")
	if err != nil {
		log.Fatal(err)
	}
	cfg = cfg1

	gs, err := startTestGrpcServer()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		log.Fatal(gs.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port)))
	}()

	logClient, err := client.NewClient(cfg.Server.Host, cfg.Server.Port)
	if err != nil {
		log.Fatal(err)
	}
	loggerClient = logClient

	os.Exit(m.Run())
}

func startTestGrpcServer() (*server.Server, error) {
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
	services := service.NewLogService(repo)
	handlers := handler.NewHandler(services)
	return server.NewServer(handlers, server.LoggingUnaryInterceptor), nil
}

func TestLog(t *testing.T) {
	type args struct {
		ctx     context.Context
		logItem *domain.LogItem
	}
	type want struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Success: send log",
			args: args{
				ctx: context.Background(),
				logItem: &domain.LogItem{
					Entity:   domain.EntityUser,
					Action:   domain.ActionCreate,
					EntityID: 1,
				},
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "Fail: Log",
			args: args{
				ctx:     context.Background(),
				logItem: &domain.LogItem{},
			},
			want: want{
				err: domain.ErrorInvalidAction,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := loggerClient.SendLogRequest(tt.args.ctx, tt.args.logItem)
			if err != tt.want.err {
				t.Errorf("Log() error = %v, wantErr %v", err, tt.want.err)
				return
			}
		})
	}
}
