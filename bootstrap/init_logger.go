package bootstrap

import (
	log "github.com/sirupsen/logrus"
	"os"
)

var Log = new(logger)

type LogField log.Fields

type logger struct {
}

// Debug 调试
func (*logger) Debug(key string, value interface{}, msg ...string) {
	log.WithField(key, value).Debug(msg)
}

// Error 错误
func (*logger) Error(key string, value interface{}, msg ...string) {
	log.WithField(key, value).Error(msg)
}

func (r *logger) _init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(parseLevel())
}

func parseLevel() log.Level {
	switch Cfg.GetVal("log.level").String() {
	case "trace":
		return log.TraceLevel
	case "debug":
		return log.DebugLevel
	case "info":
		return log.InfoLevel
	case "warn":
		return log.WarnLevel
	case "error":
		return log.ErrorLevel
	case "fatal":
		return log.FatalLevel
	}
	return log.TraceLevel
}
