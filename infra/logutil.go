package infra

import "github.com/sirupsen/logrus"

type LogUtil struct {
	Name   string
	Logger *logrus.Entry
}

func New(name string) *LogUtil {
	return &LogUtil{
		Name:   name,
		Logger: DefaultLogger.WithField("logger", name),
	}
}
func NewLogger() *logrus.Logger {
	return logrus.New()
}

var DefaultLogger = NewLogger()
