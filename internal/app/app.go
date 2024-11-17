package app

import (
	"github.com/puny-activity/metadata-parser-audio/internal/api/grpc"
	"github.com/puny-activity/metadata-parser-audio/internal/api/grpc/controller"
	"github.com/puny-activity/metadata-parser-audio/pkg/werr"
	"log/slog"
)

type Config interface {
	GetGRPCConfig() grpc.Config
}

type App struct {
	config     Config
	grpcServer *grpc.Server
	log        *slog.Logger
}

func New(config Config, logger *slog.Logger) *App {
	return &App{
		config: config,
		log:    logger,
	}
}

func (a *App) Start() error {
	a.log.Info("Launching the application")

	grpcController := controller.New(a.log)
	a.grpcServer = grpc.New(a.config.GetGRPCConfig(), grpcController, a.log)
	err := a.grpcServer.Start()
	if err != nil {
		return werr.Wrap("failed to start grpc server", err)
	}

	a.log.Info("The application is run")
	return nil
}

func (a *App) Stop() error {
	a.log.Info("Stopping the application")

	a.grpcServer.Stop()

	a.log.Info("The application is stopped")
	return nil
}
