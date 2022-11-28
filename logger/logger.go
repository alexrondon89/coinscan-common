package logger

import (
	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {
	logger := logrus.New()
	logger.Info("logrus created")
	return logger
}
