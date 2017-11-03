package main

import (
	"fmt"
	"github.com/iftsoft/gopack/lla"
)

var appLog lla.LogAgent

func InitLoggerAPP(level int) {
	appLog.Init(level, "APP")
}

func main() {
	fmt.Println("-------BEGIN------------")

	logCfg := lla.LogConfig{"../../../../log/", "test_gopack",
		lla.LogLevelTrace, lla.LogLevelTrace, 10, 1, 10000}

	lla.StartFileLogger(&logCfg)
	InitLoggerAPP(lla.LogLevelTrace)

	name := "Test Name"
	RunLoggingTest(name)

	fmt.Println("-------END------------")
	lla.StopFileLogger()
}

func RunLoggingTest(name interface{}) {
	defer appLog.PanicRecover()
	appLog.Trace("Value: %v, Type: %T", name, name)
	appLog.Dump("Value: %v, Type: %T", name, name)
	appLog.Debug("Value: %v, Type: %T", name, name)
	appLog.Info("Value: %v, Type: %T", name, name)
	appLog.Warn("Value: %v, Type: %T", name, name)
	appLog.Error("Value: %v, Type: %T", name, name)
	appLog.Panic("Value: %v, Type: %T", name, name)
}
