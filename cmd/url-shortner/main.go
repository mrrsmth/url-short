package main

import (
	"fmt"
	"log/slog"
	"os"
	"url-shortner/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	//todo clean env +
	//logger slog
	//storage sqlite
	//router chi chi render
	//run server

	cfg := config.MustLoad()

	fmt.Println(cfg)

	log := setupLogger(cfg.Env)
	log.Info("starting url-shortner",
		slog.String("env", cfg.Env),
		slog.String("storage_path", cfg.StoragePath),
		slog.String("address", cfg.HTTPServer.Address),
		slog.String("timeout", cfg.HTTPServer.Timeout.String()),
		slog.String("idleTimeout", cfg.HTTPServer.IdleTimeout.String()),
	)
	log.Debug("debug message and enabled")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default: // If env config is invalid, set prod settings by default due to security
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log

}
