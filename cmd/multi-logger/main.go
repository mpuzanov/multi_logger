package main

import (
	"context"
	"log"
	"multi_logger/internal/config"
	"multi_logger/internal/glogger"
	"multi_logger/internal/services"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	logger := glogger.BuildLogger(cfg.LoggerName, cfg.Level)

	logger.Debugf("%#v", cfg)

	// кидаем логер в контекст
	ctx := glogger.ContextWithLogger(context.Background(), logger)

	// создаем сервис с бизнес-логикой программы
	s := services.New(ctx, logger)

	if err := s.Run(); err != nil {
		logger.Fatal(err)
	}

}
