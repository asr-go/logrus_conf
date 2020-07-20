package logrusconf

import (
	"testing"

	"github.com/sirupsen/logrus"
)

// TestLog 测试日志
func TestLog(t *testing.T) {
	Init(logrus.DebugLevel, "./logs/logrus.log")
	for index := 0; index < 10; index++ {
		logrus.WithField("Success", "True").WithField("Failed", "False").Debug("DEBU")
	}
	for index := 0; index < 10; index++ {
		logrus.WithField("Success", "True").WithField("Failed", "False").Info("INFO")
	}
	for index := 0; index < 10; index++ {
		logrus.WithField("Success", "True").WithField("Failed", "False").Warn("WARN")
	}
	for index := 0; index < 10; index++ {
		logrus.WithField("Success", "True").WithField("Failed", "False").Error("ERRO")
	}
}
