// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/ZQCard/kbk-layout/internal/biz"
	"github.com/ZQCard/kbk-layout/internal/conf"
	"github.com/ZQCard/kbk-layout/internal/data"
	"github.com/ZQCard/kbk-layout/internal/server"
	"github.com/ZQCard/kbk-layout/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/sdk/trace"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(env *conf.Env, confServer *conf.Server, registry *conf.Registry, confData *conf.Data, bootstrap *conf.Bootstrap, logger log.Logger, tracerProvider *trace.TracerProvider) (*kratos.App, func(), error) {
	db := data.NewMysqlCmd(bootstrap, logger)
	client := data.NewRedisClient(confData)
	dataData, cleanup, err := data.NewData(bootstrap, db, client, logger)
	if err != nil {
		return nil, nil, err
	}
	exampleRepo := data.NewExampleRepo(dataData, logger)
	exampleUsecase := biz.NewExampleUsecase(exampleRepo, logger)
	exampleService := service.NewExampleService(exampleUsecase, logger)
	grpcServer := server.NewGRPCServer(confServer, exampleService, logger)
	httpServer := server.NewHTTPServer(bootstrap, confServer, exampleService, logger)
	registrar := data.NewRegistrar(registry)
	app := newApp(logger, grpcServer, httpServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
