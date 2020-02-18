package trunk

import (
	"log"
	"os"
)

var logDebug *log.Logger
var logErr *log.Logger
var logFatal *log.Logger
var logInfo *log.Logger
var logSuccess *log.Logger
var logWarn *log.Logger

func init() {
	ansiGreen := "\u001b[32m"
	ansiBlue := "\u001b[34m"
	ansiRed := "\u001b[31m"
	ansiReset := "\u001b[0m"
	ansiWhite := "\u001b[37;1m"
	ansiYellow := "\u001b[33m"

	logDebug = log.New(os.Stdout, ansiGreen+"[debug] "+ansiReset, 0)
	logErr = log.New(os.Stderr, ansiRed+"[error] "+ansiReset, 0)
	logFatal = log.New(os.Stderr, ansiRed+"[fatal] "+ansiReset, 0)
	logInfo = log.New(os.Stdout, ansiWhite+"[info] "+ansiReset, 0)
	logSuccess = log.New(os.Stdout, ansiBlue+"[success] "+ansiReset, 0)
	logWarn = log.New(os.Stdout, ansiYellow+"[warn] "+ansiReset, 0)
}

// LogDebug will log a message as debug
func LogDebug(message string) {
	logDebug.Println(message)
}

// LogDebugRaw will log an error's message as debug
func LogDebugRaw(message error) {
	LogDebug(message.Error())
}

// LogErr will log a message at an "error" level
func LogErr(message string) {
	logErr.Println(message)
}

// LogErrRaw will log an error's message at an "error" level
func LogErrRaw(message error) {
	LogErr(message.Error())
}

// LogFatal will log a message as fatal and exit
func LogFatal(message string) {
	logFatal.Fatalln(message)
}

// LogFatalErr will log an error's message as fatal and exit
func LogFatalErr(message error) {
	LogFatal(message.Error())
}

// LogInfo will log a message at an "info" level
func LogInfo(message string) {
	logInfo.Println(message)
}

// LogSuccess will log a successful action message at an "info" level
func LogSuccess(message string) {
	logSuccess.Println(message)
}

// LogWarn will log a warning message at an "info" level
func LogWarn(message string) {
	logWarn.Println(message)
}

// LogWarnErr will log an error's message as a warning at an "info" level
func LogWarnErr(message error) {
	LogWarn(message.Error())
}
