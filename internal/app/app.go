package app

import (
	"log/slog"
)

type App struct {
	log *slog.Logger
}

func New(logger *slog.Logger) *App {
	return &App{
		log: logger,
	}
}

func (a *App) Start() error {
	a.log.Info("Launching the application")

	a.log.Info("The application is run")
	return nil
}

func (a *App) Stop() error {
	a.log.Info("Stopping the application")

	a.log.Info("The application is stopped")
	return nil
}
