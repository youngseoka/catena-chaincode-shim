package logger

import (
	"github.com/sirupsen/logrus"
)

func SetBasicLogger(moduleName string, logLevel string) {
	logger := logrus.WithFields(logrus.Fields{"module": moduleName})
	logrus.SetFormatter(&logrus.JSONFormatter{})
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logger.Warn("Wrong logger level. Set with debug.")
		level = logrus.DebugLevel
	}
	logrus.SetLevel(level)
}
