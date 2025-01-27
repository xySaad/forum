package log

import (
	"fmt"
	"os"
	"time"
)

var logFile *os.File

func Initlog() error {
	err := os.MkdirAll("./logs", 0o777)
	if err != nil {
		return err
	}
	logFile, err = os.Create("./logs/" + time.Now().Format(time.DateOnly) + ".log")
	if err != nil {
		return err
	}
	return nil
}

func Info(v ...any) {
	err := Log("INFO", v)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error loging error: ", v)
	}
}

func Warn(v ...any) {
	err := Log("WARN", v)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error loging error: ", v)
	}
}

func Error(v ...any) {
	err := Log("ERROR", v)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error loging error: ", v)
	}
}

func Log(logLevel string, msg any) error {
	_, err := fmt.Fprintln(logFile, fmt.Sprintf("%v %v: %v", time.Now().Format(time.DateTime), logLevel, msg))
	return err
}
