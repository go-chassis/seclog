package main

import (
	"github.com/go-chassis/openlog"
	"github.com/go-chassis/paas-lager"
	"github.com/go-chassis/paas-lager/rotate"
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

	logger.Debug("check-info", openlog.WithTags(openlog.Tags{
		"info": "something",
	}))

	logger.Warn("failed-to-do-something", openlog.WithTags(openlog.Tags{
		"info": "something",
	}))
	logger.Warnf("failed-to-do-%s-something", "1")

	logger.Error("failed-to-do-something")

	logger.Info("shutting-down")

}
