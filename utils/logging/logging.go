package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func NewLogger(_ string) *logrus.Logger {
	if Logger != nil {
		return Logger
	}
	file, err := os.OpenFile("./log/debug.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0774)
	if err != nil {
		Logger.Fatal(err)
	}
	Logger = &logrus.Logger{
		Out:          file,
		Formatter:    &logrus.JSONFormatter{},
		Hooks:        make(logrus.LevelHooks),
		Level:        logrus.InfoLevel,
		ReportCaller: true,
	}
	return Logger
}
