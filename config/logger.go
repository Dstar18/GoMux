package config

import (
	"os"
	"time"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func InitLogger() {
	Logger = logrus.New()

	// open file
	logFile := "logs/server.log"
	file, err := os.OpenFile(logFile, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("Failed to open log file: %v", err)
	}

	// config logrus
	Logger.SetOutput(file)
	Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})
	Logger.SetLevel(logrus.InfoLevel)

	// Hook for save log to file
	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  logFile,
		logrus.WarnLevel:  logFile,
		logrus.ErrorLevel: logFile,
	}

	Logger.AddHook(lfshook.NewHook(pathMap, &logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	}))

	Logger.Info("Logger initialized successfully!")
}
