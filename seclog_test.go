package seclog_test

import (
	"fmt"
	"github.com/go-chassis/seclog"
	"testing"
)

func TestNewLogger(t *testing.T) {
	seclog.Init(seclog.Config{
		LoggerLevel:   "DEBUG",
		LogFormatText: true,
		Writers:       []string{"stdout"},
	})

	logger := seclog.NewLogger("example")
	logger.Debug("hi")
	logger.Info("hi")
	logger.Warn("hi")
}
func BenchmarkInit(b *testing.B) {
	seclog.Init(seclog.Config{
		LoggerLevel:   "DEBUG",
		LogFormatText: true,
		Writers:       []string{"stdout"},
	})

	logger := seclog.NewLogger("example")

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("test")
		}
	})
	b.ReportAllocs()
	logger.Debug("hi")
}

type w struct {
}

func (w *w) Write(p []byte) (n int, err error) {
	fmt.Print("fake")
	return 2, nil
}
func TestRegisterWriter(t *testing.T) {
	seclog.RegisterWriter("test", &w{})
}
