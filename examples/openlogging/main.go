package main

import (
	"github.com/go-chassis/paas-lager"
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
	openlogging.SetLogger(logger)

	openlogging.Debug("check-info", openlogging.WithTags(openlogging.Tags{
		"info": "something",
	}))

	openlogging.Warn("failed-to-do-somthing", openlogging.WithTags(openlogging.Tags{
		"info": "something",
	}))

	openlogging.Error("failed-to-do-somthing")

	openlogging.Info("shutting-down")

}
