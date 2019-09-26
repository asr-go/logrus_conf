package config

import (
	"testing"

	"github.com/sirupsen/logrus"
)

// TestLog 测试日志
func TestLog(t *testing.T) {
	Init("./logs/logrus.log")
	for index := 0; index < 10; index++ {
		logrus.WithField("Success", "True").WithField("Failed", "False").Debug("DEBU")
		logrus.WithField("Success", "True").WithField("Failed", "False").Info("INFO")
		logrus.WithField("Success", "True").WithField("Failed", "False").Warn("WARN")
		logrus.WithField("Success", "True").WithField("Failed", "False").Error("ERRO")
	}
}
