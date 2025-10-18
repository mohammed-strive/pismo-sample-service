package utils

import (
	"context"
	"pismo-service/middlewares"

	"github.com/sirupsen/logrus"
)

func GetLogger(c context.Context) *logrus.Entry {
	logger := c.Value(middlewares.LoggerKey)
	if entry, ok := logger.(*logrus.Entry); ok {
		return entry
	}

	return logrus.New().WithFields(logrus.Fields{"context_logger_missing": true})
}
