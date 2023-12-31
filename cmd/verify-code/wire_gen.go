// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"verify-code/internal/biz"
	"verify-code/internal/conf"
	"verify-code/internal/data"
	"verify-code/internal/server"
	"verify-code/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	greeterRepo := data.NewGreeterRepo(dataData, logger)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, logger)
	greeterService := service.NewGreeterService(greeterUsecase)
	studentRepo := data.NewStudentRepo(dataData, logger)
	studentUsecase := biz.NewStudentUsecase(studentRepo, logger)
	studentService := service.NewStudentService(studentUsecase)
	grpcServer := server.NewGRPCServer(confServer, greeterService, studentService, logger)
	httpServer := server.NewHTTPServer(confServer, greeterService, studentService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
