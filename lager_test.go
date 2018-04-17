package log_test

import (
	"fmt"
	"github.com/ServiceComb/paas-lager"
	"testing"
)

func BenchmarkInit(b *testing.B) {
	log.Init(log.Config{
		LoggerLevel:   "DEBUG",
		LogFormatText: true,
		Writers:       []string{"file", "stdout"},
	})

	logger := log.NewLogger("example")

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("test")
		}
	})
	b.ReportAllocs()
}

type w struct {
}

func (w *w) Write(p []byte) (n int, err error) {
	fmt.Print("fake")
	return 2, nil
}
func TestRegisterWriter(t *testing.T) {
	log.RegisterWriter("test", &w{})
}
