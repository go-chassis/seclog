package main

import (
	"github.com/go-chassis/openlog"
	"github.com/go-chassis/seclog"
)

func main() {
	seclog.Init(seclog.Config{
		LoggerLevel:   "DEBUG",
		LoggerFile:    "test.log",
		EnableRsyslog: false,
		LogFormatText: true,
		Writers:       []string{"file", "stdout"},
	})

	logger := seclog.NewLogger("example")
	openlog.SetLogger(logger)

	openlog.Debug("check-info", openlog.WithTags(openlog.Tags{
		"info": "something",
	}))

	openlog.Warn("failed-to-do-somthing", openlog.WithTags(openlog.Tags{
		"info": "something",
	}))

	openlog.Error("failed-to-do-somthing")

	openlog.Info("shutting-down")

}
