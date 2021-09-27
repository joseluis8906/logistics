package log

import (
	log "github.com/sirupsen/logrus"
)

type (
	CustomLogger struct {
		Env       string
		Formatter log.Formatter
	}
)

func (l CustomLogger) Format(entry *log.Entry) ([]byte, error) {
	entry.Data["env"] = l.Env
	return l.Formatter.Format(entry)
}

func SetLogLevel(levelStr string) {
	switch levelStr {
	case "panic":
		log.SetLevel(log.PanicLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "warning":
		log.SetLevel(log.WarnLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "trace":
		log.SetLevel(log.TraceLevel)
	default:
		log.SetLevel(log.WarnLevel)
	}
}
