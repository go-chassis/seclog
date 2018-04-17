package main

import (
	"fmt"

	"github.com/ServiceComb/paas-lager"
	"github.com/ServiceComb/paas-lager/rotate"
	"github.com/ServiceComb/paas-lager/third_party/forked/cloudfoundry/lager"
)

func main() {
	log.Init(log.Config{
		LoggerLevel:   "DEBUG",
		LoggerFile:    "test.log",
		EnableRsyslog: false,
		LogFormatText: false,
		Writers:       []string{"file", "stdout"},
	})

	logger := log.NewLogger("example")
	rotate.RunLogRotate("test.log", &rotate.RotateConfig{}, logger)
	logger.Infof("Hi %s, system is starting up ...", "paas-bot")

	logger.Debug("check-info", lager.Data{
		"info": "something",
	})

	err := fmt.Errorf("Oops, error occurred")
	logger.Warn("failed-to-do-somthing", err, lager.Data{
		"info": "something",
	})

	err = fmt.Errorf("This is an error")
	logger.Error("failed-to-do-somthing", err)

	logger.Info("shutting-down")

}
