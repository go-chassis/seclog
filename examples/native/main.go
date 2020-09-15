package main

import (
	"github.com/go-chassis/openlog"
	"github.com/go-chassis/seclog"
)

func main() {
	seclog.Init(seclog.Config{
		LoggerLevel:   "DEBUG",
		LoggerFile:    "test.log",
		LogFormatText: false,
		Writers:       []string{"file", "stdout"},
	})

	logger := seclog.NewLogger("example")

	logger.Debug("check-info", openlog.WithTags(openlog.Tags{
		"info": "something",
	}))

	logger.Warn("failed-to-do-something", openlog.WithTags(openlog.Tags{
		"info": "something",
	}))

	logger.Error("failed-to-do-something")

	logger.Info("shutting-down")

}
