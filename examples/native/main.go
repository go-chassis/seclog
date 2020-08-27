package main

import (
	"github.com/go-chassis/openlog"
	"github.com/go-chassis/seclog"
	"github.com/go-chassis/seclog/rotate"
)

func main() {
	seclog.Init(seclog.Config{
		LoggerLevel:   "DEBUG",
		LoggerFile:    "test.log",
		EnableRsyslog: false,
		LogFormatText: false,
		Writers:       []string{"file", "stdout"},
	})

	logger := seclog.NewLogger("example")
	rotate.RunLogRotate("test.log", &rotate.RotateConfig{}, logger)

	logger.Debug("check-info", openlog.WithTags(openlog.Tags{
		"info": "something",
	}))

	logger.Warn("failed-to-do-something", openlog.WithTags(openlog.Tags{
		"info": "something",
	}))

	logger.Error("failed-to-do-something")

	logger.Info("shutting-down")

}
