package log

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var SysLog = logrus.New()

func init() {
	logLevelStr := os.Getenv("LOG_LEVEL")
	if logLevelStr == "" {
		logLevelStr = "INFO"
	}
	logLevel, err := logrus.ParseLevel(logLevelStr)
	if err != nil {
		panic(fmt.Errorf("error loading logLevel: %s", err))
	}
	fmt.Printf("use log with level %s\n", logLevelStr)
	SysLog.SetLevel(logLevel)
	SysLog.SetOutput(os.Stdout)
	SysLog.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}
