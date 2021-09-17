package lager

import (
	"bytes"
	"github.com/go-chassis/seclog/third_party/forked/cloudfoundry/lager/color"
	"io"
	"sync"
)

const logBufferSize = 1024

const (
	// ColorModeAuto only print colorful log when writer is stdout
	ColorModeAuto = "auto"
	// ColorModeNever never print colorful log
	ColorModeNever = "never"
	// ColorModeAlways always print colorful log
	ColorModeAlways = "always"
)

// A Sink represents a write destination for a Logger. It provides
// a thread-safe interface for writing logs
type Sink interface {
	//Log to the sink.  Best effort -- no need to worry about errors.
	Log(level LogLevel, payload []byte)
}

type writerSink struct {
	writer      io.Writer
	minLogLevel LogLevel
	name        string
	writeL      *sync.Mutex
	colorMode   string
}

//NewWriterSink is function which returns new struct object
func NewWriterSink(name string, writer io.Writer, minLogLevel LogLevel, colorMode string) Sink {
	return &writerSink{
		writer:      writer,
		minLogLevel: minLogLevel,
		writeL:      new(sync.Mutex),
		name:        name,
		colorMode:   colorMode,
	}
}

func (sink *writerSink) Log(level LogLevel, log []byte) {
	if level < sink.minLogLevel {
		return
	}
	if sink.colorMode == ColorModeAlways ||
		(sink.colorMode == ColorModeAuto && sink.name == "stdout") {
		switch level {
		case FATAL:
			log = bytes.Replace(log, []byte("FATAL"), color.FatalByte, -1)
		case ERROR:
			log = bytes.Replace(log, []byte("ERROR"), color.ErrorByte, -1)
		case WARN:
			log = bytes.Replace(log, []byte("WARN"), color.WarnByte, -1)
		case INFO:
			log = bytes.Replace(log, []byte("INFO"), color.InfoByte, -1)
		default:
		}
	}
	log = append(log, '\n')
	sink.writeL.Lock()
	sink.writer.Write(log)
	sink.writeL.Unlock()
}
