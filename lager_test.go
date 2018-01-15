package stlager_test

import (
	"github.com/ServiceComb/paas-lager"
	"testing"
)

func BenchmarkInit(b *testing.B) {
	stlager.Init(stlager.Config{
		LoggerLevel:   "DEBUG",
		LogFormatText: true,
		Writers:       []string{"file"},
	})

	logger := stlager.NewLogger("example")

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("test")
		}
	})
	b.ReportAllocs()
}
