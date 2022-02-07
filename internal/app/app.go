package app

import (
	"fmt"
	v1 "github.com/swanden/service-template/internal/controller/http/v1"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"github.com/swanden/service-template/pkg/config"
	"github.com/swanden/service-template/pkg/httpserver"
	"github.com/swanden/service-template/pkg/logger"
)

func Run(configFile string) {
	conf, err := config.New(configFile)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	var l logger.Interface
	l, err = logger.New(conf.Logger.File, conf.Logger.Level)
	if err != nil {
		log.Fatalf("Logger error: %s", err)
	}

	handler := gin.New()
	v1.NewRouter(handler)
	httpServer := httpserver.New(handler, httpserver.Port(conf.HTTP.Port))

	l.Info("app - Run - server start on " + conf.HTTP.Port)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
