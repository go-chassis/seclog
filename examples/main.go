package main

import (
	"fmt"

	"github.com/ServiceComb/paas-lager"
	"github.com/ServiceComb/paas-lager/third_party/forked/cloudfoundry/lager"
)

func main() {
	stlager.Init(stlager.Config{
		LoggerLevel:   "DEBUG",
		LoggerFile:    "",
		EnableRsyslog: false,
		LogFormatText: false,
	})

	logger := stlager.NewLogger("example")

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
