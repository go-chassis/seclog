# seclog

A secure log tool fo go
### Usage
Create logger
```
seclog.Init(seclog.Config{
        LoggerLevel:   loggerLevel,
        LoggerFile:    loggerFile,
        LogFormatText:  false,
})

logger := seclog.NewLogger(component)
```

* LoggerLevel: 日志级别由低到高分别为 DEBUG, INFO, WARN, ERROR, FATAL 共5个级别，这里设置的级别是日志输出的最低级别，只有不低于该级别的日志才会输出
* LoggerFile: 输出日志的文件名，为空则输出到 os.Stdout
* LogFormatText: 设定日志的输出格式是 json 还是 plaintext

Create logger with multiple sinker
```go
	seclog.Init(seclog.Config{
		LoggerLevel:   "DEBUG",
		LoggerFile:    "test.log",
		LogFormatText: false,
		Writers:       []string{"file", "stdout"},
	})

	logger := seclog.NewLogger("example")
```

Custom your own sinker
```go
type w struct {
}

func (w *w) Write(p []byte) (n int, err error) {
	fmt.Print("fake")
	return 0, nil
}
func main() {
	seclog.RegisterWriter("test", &w{})
}

```
See [Examples](examples)

Change log level on-fly
```go
l.SetLogLevel(lager.ERROR)
```
