package services

import (
	"context"
	"multi_logger/internal/glogger"
)

// App ...
type App struct {
	ctx context.Context
	log glogger.Logger
}

// New ...
func New(ctx context.Context, logger glogger.Logger) *App {
	return &App{ctx: ctx, log: logger}
}

// Run запуск сервиса
func (app *App) Run() error {

	app.log.Info("Выполняем обработку")

	app.Foo()

	app.log.Info("Закончили")
	return nil
}

// Foo ...
func (app *App) Foo() {
	logger := glogger.LoggerFromContext(app.ctx)

	logger.Debugf("вызов функции Foo() : %s", "(logger from context)")
}
