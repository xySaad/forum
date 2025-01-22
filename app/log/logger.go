package logs

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var (
	Logger  *log.Logger
	logFile *os.File
)

func InitLogger() error {
	logFolder := "./app/log/logs"
	var err error
	timestamp := time.Now().Format("20060102_150405")
	logFileName := fmt.Sprintf("%s/server_%s.log", logFolder, timestamp)
	logFile, err = os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}

	multiWriter := io.MultiWriter(os.Stdout, logFile)

	Logger = log.New(multiWriter, "", log.LstdFlags)
	Logger.Println("logger created Succefly")
	return nil
}

func CloseLogger() {
	if logFile != nil {
		logFile.Close()
	}
}
