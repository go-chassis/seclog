# Logging Utility for Go-Chassis

A structured logger for Go

#### Usage
Create logger (complex mode)
```
log.Init(paas_lager.Config{
        LoggerLevel:   loggerLevel,
        LoggerFile:    loggerFile,
        LogFormatText:  false,
})

logger := paas_lager.NewLogger(component)
```

* LoggerLevel: 日志级别由低到高分别为 DEBUG, INFO, WARN, ERROR, FATAL 共5个级别，这里设置的级别是日志输出的最低级别，只有不低于该级别的日志才会输出。所以，如果不想打印 WARN 以下级别的日志，可以把 LoggerLevel 设置为 WARN。
从减小日志大小的考虑，建议在单元测试的时候，可以把 LoggerLevel 级别设置的比较低，比如 DEBUG，而在正式部署的系统代码中，建议级别设置为 WARN，以减少日志频繁输出带来的高 IO 对 VM 和 PaaS 日志后台的存储和处理的负担。
* LoggerFile: 输出日志的文件名，为空则输出到 os.Stdout。建议路径统一放置在 /var/paas/sys/log 目录下，各个模块分别创建各自的子目录，这样后续可以通过 logrotate 统一进行日志的 rotate 防爆处理。
* LogFormatText: 设定日志的输出格式是 json 还是 plaintext (类似于log4j)，默认为 false，不建议修改，如果开发过程中想本地查看日志的话，可以设定 LoggerFile 和 LogFormatText 为 true，这样会输出类似于 log4j 格式的本地日志。但是建议正式的生产环境还是用 json 的格式输出。

对于大多数的使用场景，可以只配置下面的参数即可。
Create logger (simple mode)
```go
	log.Init(log.Config{
		LoggerLevel:   "DEBUG",
		LoggerFile:    "test.log",
		EnableRsyslog: false,
		LogFormatText: false,
		Writers:       []string{"file", "stdout"},
	})

	logger := log.NewLogger("example")
```
Run log rotate
```go
rotate.RunLogRotate("test.log", &rotate.RotateConfig{}, logger)
```

See [Example](examples/main.go)