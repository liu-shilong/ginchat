package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func createLogFile() (*os.File, error) {
	now := time.Now()
	date := now.Format("2006-01-02")
	filename := fmt.Sprintf("runtime/logs/%s.log", date)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func SetupLogger() (*logrus.Logger, error) {
	logger := logrus.New()
	file, err := createLogFile()
	if err != nil {
		return nil, err
	}
	logger.Out = file
	logger.SetFormatter(&logrus.JSONFormatter{})
	return logger, nil
}

func Record(msg interface{}) {
	logrus.Info(msg)
}
