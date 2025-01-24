package config

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var (
	Logger      *log.Logger
	MultiLogger *log.Logger
	logFile     *os.File
)

func InitLogger() error {
	logFolder := "./logs"
	err := os.Mkdir(logFolder, 0o766)
	if err != nil && !os.IsExist(err) {
		return err
	}
	timestamp := time.Now().Format(time.DateOnly)
	logFileName := fmt.Sprintf("%s/server_%s.log", logFolder, timestamp)
	logFile, err = os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0o644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}
	MultiLogger = log.New(io.MultiWriter(os.Stdout, logFile), "", log.LstdFlags)
	Logger = log.New(logFile, "", log.LstdFlags)
	return nil
}

func CloseLogger() {
	if logFile != nil {
		logFile.Close()
	}
}
