package app

import (
	"github.com/swanden/service-template/pkg/config"
	"github.com/swanden/service-template/pkg/logger"
	"log"
)

func Run(configFile string) {
	conf, err := config.New(configFile)
	if err != nil {
		log.Fatal(err)
	}

	var l logger.Interface
	l, err = logger.New(conf.Logger.File, conf.Logger.Level)
	if err != nil {
		log.Fatal(err)
	}

	l.Debug("application started", logger.Filed{Name: "TestName", Value: "TestValue"})
}
