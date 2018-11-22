package main

import (
	"github.com/go-chassis/paas-lager"
	"github.com/go-chassis/paas-lager/rotate"
	"github.com/go-mesh/openlogging"
)

func main() {
	log.Init(log.Config{
		LoggerLevel:   "DEBUG",
		LoggerFile:    "test.log",
		EnableRsyslog: false,
		LogFormatText: true,
		Writers:       []string{"file", "stdout"},
	})

	logger := log.NewLogger("example")
	rotate.RunLogRotate("test.log", &rotate.RotateConfig{}, logger)
	logger.Infof("Hi %s, system is starting up ...", "paas-bot")

	logger.Debug("check-info", openlogging.WithTags(openlogging.Tags{
		"info": "something",
	}))

	logger.Warn("failed-to-do-somthing", openlogging.WithTags(openlogging.Tags{
		"info": "something",
	}))
	logger.Warnf("failed-to-do-%s-somthing", "1")

	logger.Error("failed-to-do-somthing")

	logger.Info("shutting-down")

}
