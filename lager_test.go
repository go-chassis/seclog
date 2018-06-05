package log_test

import (
	"fmt"
	"github.com/ServiceComb/paas-lager"
	"gopkg.in/yaml.v2"
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

func TestInit(t *testing.T) {
	b := []byte(`
writers: [file, stdout]

`)
	c := &log.Config{}
	err := yaml.Unmarshal(b, c)
	if err != nil {
		t.Error(err)
	}
	if len(c.Writers) != 2 {
		t.Error(c.Writers)
	}
}
