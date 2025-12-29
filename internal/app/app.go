package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/Kulichanin/Authorization-Service/internal/app/grpc"
	"google.golang.org/grpc"
)

type App struct {
	GRPCSrc *grpc.Server
}

func NewApp(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	// TODO: инициализировать приложение (storage)

	// TODO: инициализировать auth service (auth)

	grpcAPP := grpcapp.NewApp(log, grpcPort)

	return &App{
		GRPCSrc: grpcApp,
	}
}
