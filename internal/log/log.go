package log

import (
	"errors"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

func NewLogger(output, format, level string) (*logrus.Logger, error) {
	logPath, err := getOutputPath(output)
	if err != nil {
		return nil, err
	}

	logLevel := getLevel(level)
	logFormat := getFormat(format)

	logger := &logrus.Logger{
		Out:       logPath,
		Formatter: logFormat,
		Level:     logLevel,
	}
	return logger, nil
}

func getOutputPath(dir string) (*os.File, error) {
	_, err := os.ReadDir(dir)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, os.ErrNotExist
		}
		return nil, err
	}

	filename := fmt.Sprintf("log_%s.log", time.Now().Format("02_01_2006"))
	basePath := path.Base(dir)
	logFile := path.Join(basePath, filename)

	logReader, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	return logReader, nil
}

func getLevel(level string) logrus.Level {
	var logLevel logrus.Level

	switch level {
	case "debug":
		logLevel = logrus.DebugLevel
	case "info":
		logLevel = logrus.InfoLevel
	case "warn":
		logLevel = logrus.WarnLevel
	case "error":
		logLevel = logrus.ErrorLevel
	default:
		logLevel = logrus.InfoLevel
	}

	return logLevel
}

func getFormat(format string) logrus.Formatter {
	var logFormat logrus.Formatter

	switch format {
	case "json":
		logFormat = &logrus.JSONFormatter{}
	case "text":
		logFormat = &logrus.TextFormatter{}
	default:
		logFormat = &logrus.JSONFormatter{}
	}

	return logFormat
}
