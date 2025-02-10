package log

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"time"
)

var logFile io.Writer
var logFilename = "./logs/" + time.Now().Format(time.DateOnly) + ".log"

func Initlog() error {
	err := os.MkdirAll("./logs", 0o777)
	if err != nil {
		return err
	}
	logFile, err = os.OpenFile(logFilename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	return nil
}

func Info(v ...any) {
	Log("INFO", true, v...)
}

func Warn(v ...any) {
	Log("WARN", false, v...)
}

func Error(v ...any) {
	Log("ERROR", false, v...)
}

func Debug(v ...any) {
	// TODO: Invoke only in debug mode
	Log("DEBUG", true, v...)
}

func Fatal(v ...any) {
	Log("FATAL", true, v...)
	os.Exit(1)
}

func Log(logLevel string, isMulti bool, msg ...any) {
	var source string
	logger := logFile
	if isMulti {
		logger = io.MultiWriter(logFile, os.Stdout)
	}
	if logLevel == "DEBUG" || !isMulti {
		_, file, line, _ := runtime.Caller(2)
		source = fmt.Sprintf("source %v:%v\n", file, line)
	}
	_, err := fmt.Fprintf(logger, "%v [%v] %v\n%v", time.Now().Format(time.RFC3339), logLevel, msg, source)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error logging error:", err)
		return
	}
}
