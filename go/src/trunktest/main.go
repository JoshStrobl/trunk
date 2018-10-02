package main

import (
	"errors"
	"trunk"
)

func main() {
	logErr := errors.New("this is an Error passed to LogErrRaw")
	logWarn := errors.New("this is an Error passed to LogWarnErr")

	trunk.LogInfo("running through various trunk functions")
	trunk.LogErr("this is showing LogErr in action")
	trunk.LogErrRaw(logErr)
	trunk.LogSuccess("this is showing LogSuccess in action")
	trunk.LogWarn("this is showing LogWarn in action")
	trunk.LogWarnErr(logWarn)
	trunk.LogFatal("this is showing LogFatal in action. Exiting")
}
