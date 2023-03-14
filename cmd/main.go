package main

import (
	"IntegrationLab1/configs"
	"IntegrationLab1/internal/delivery/http/v1"
	"IntegrationLab1/internal/service"
	"IntegrationLab1/pkg/httpServer"
	"IntegrationLab1/pkg/logger"
	"context"
	"github.com/go-playground/validator/v10"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log := logger.NewLogger()
	log.Info("logger has been started")

	cfg := configs.NewConfig()
	log.Info("config has been init")

	validate := validator.New()

	serv := service.NewService(cfg)
	handler := v1.NewHandler(serv, log, validate)

	server := httpServer.NewServer(handler.InitRoutes(), httpServer.Port(cfg.Server.HttpPort))

	log.Infof("server started on: [http://localhost:%s]", cfg.Server.HttpPort)

	if err := server.Run(); err != nil {
		log.Fatalf("error http server, %s", err.Error())
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := server.Shutdown(context.Background()); err != nil {
		log.Errorf("error on server shutting down, %s", err.Error())
	}
}
