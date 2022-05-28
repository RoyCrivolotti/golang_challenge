package utils

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"
	"os"
	"path/filepath"
	"testing"
	"time"
)

var Logger *zap.Logger

func Init() error {
	//Create config for the zap encoder using the ISO time format
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	//Create a file encoder that follows the standard JSON encoding
	fileEncoder := zapcore.NewJSONEncoder(config)
	consoleEncoder := zapcore.NewConsoleEncoder(config)

	//Create tmp folder if it doesn't exist
	logFilesPath := filepath.Join("./", "tmp")
	if _, err := os.Stat(logFilesPath); errors.Is(err, os.ErrNotExist) {
		if err := os.Mkdir(logFilesPath, os.ModePerm); err != nil {
			return err
		}
	}

	//Open up the log file in the Create and Append mode
	newLogFilePath := fmt.Sprintf("%s/log_%s.json", logFilesPath, time.Now().UTC().Format("2006_01_02T15_04_05"))
	logFile, err := os.OpenFile(newLogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	//Create a file writer for zap to write the messages to the log file
	writer := zapcore.AddSync(logFile)

	//Set the default log level and add the file writer, encoder and default log level to the zap core instance
	defaultLogLevel := zapcore.DebugLevel
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)

	//Create a new Zap instance with the configs created
	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.PanicLevel))

	return nil
}

func InitTest(t *testing.T) {
	Logger = zaptest.NewLogger(t)
}
