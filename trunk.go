package trunk

import (
	"log"
	"os"
	"strings"
)

var (
	LoggerDebug, LoggerErr, LoggerFatal, LoggerInfo, LoggerSuccess, LoggerWarn *log.Logger
	LevelSuccess, LevelInfo, LevelDebug, LevelWarn, LevelErr, LevelFatal       = 0, 1, 2, 3, 4, 5
)

func init() {
	ansiGreen := "\u001b[32m"
	ansiBlue := "\u001b[34m"
	ansiRed := "\u001b[31m"
	ansiReset := "\u001b[0m"
	ansiWhite := "\u001b[37;1m"
	ansiYellow := "\u001b[33m"

	LoggerDebug = log.New(os.Stdout, ansiGreen+"[debug] "+ansiReset, 0)
	LoggerErr = log.New(os.Stderr, ansiRed+"[error] "+ansiReset, 0)
	LoggerFatal = log.New(os.Stderr, ansiRed+"[fatal] "+ansiReset, 0)
	LoggerInfo = log.New(os.Stdout, ansiWhite+"[info] "+ansiReset, 0)
	LoggerSuccess = log.New(os.Stdout, ansiBlue+"[success] "+ansiReset, 0)
	LoggerWarn = log.New(os.Stdout, ansiYellow+"[warn] "+ansiReset, 0)
}

// LogLevel will log using a specific level our message and provided values
func LogLevel(level int, message string, vals ...interface{}) {
	if !strings.HasSuffix(message, "\n") { // Doesn't end with a newline
		message += "\n" // Add one
	}

	switch level {
	case LevelInfo:
		LoggerInfo.Printf(message, vals)
	case LevelDebug:
		LoggerDebug.Printf(message, vals)
	case LevelWarn:
		LoggerWarn.Printf(message, vals)
	case LevelErr:
		LoggerErr.Printf(message, vals)
	case LevelFatal:
		LoggerFatal.Fatalf(message, vals)
	default: // Success
		LoggerSuccess.Printf(message, vals)
	}
}

// LogDebug will log a message as debug
func LogDebug(message string, vals ...interface{}) {
	LogLevel(LevelDebug, message, vals)
}

// LogErr will log a message at an "error" level
func LogErr(message string, vals ...interface{}) {
	LogLevel(LevelErr, message, vals)
}

// LogFatal will log a message as fatal and exit
func LogFatal(message string, vals ...interface{}) {
	LogLevel(LevelFatal, message, vals)
}

// LogInfo will log a message at an "info" level
func LogInfo(message string, vals ...interface{}) {
	LogLevel(LevelInfo, message, vals)
}

// LogSuccess will log a successful action message at an "info" level
func LogSuccess(message string, vals ...interface{}) {
	LogLevel(LevelSuccess, message, vals)
}

// LogWarn will log a warning message at an "info" level
func LogWarn(message string, vals ...interface{}) {
	LogLevel(LevelWarn, message, vals)
}
