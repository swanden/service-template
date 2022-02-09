package app

import (
	"flag"
	"fmt"
	v1 "github.com/swanden/service-template/internal/controller/http/v1"
	"github.com/swanden/service-template/internal/domain/user/entity"
	"github.com/swanden/service-template/pkg/database"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"github.com/swanden/service-template/pkg/config"
	"github.com/swanden/service-template/pkg/httpserver"
	"github.com/swanden/service-template/pkg/logger"
)

var migrate = flag.Bool("m", false, "Run migrations")

func Run(configFile string) {
	flag.Parse()

	conf, err := config.New(configFile)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	var l logger.Interface
	l, err = logger.New(conf.Logger.File, conf.Logger.Level)
	if err != nil {
		log.Fatalf("Logger error: %s", err)
	}

	db, err := database.New(conf.Postgres.DSN)
	if err != nil {
		log.Fatalf("Database error: %s", err)
	}
	if *migrate {
		err = db.Migrate(&entity.User{})
		if err != nil {
			log.Fatalf("Migrate error: %s", err)
		}
		l.Info("app - Migrations applied")
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
