package app

import (
	grpcapp "sso/internal/app/grpc"
	"time"
)

type App struct {
	gRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) * App {
	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		gRPCSrv: grpcApp,
	}
}