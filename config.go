package logrusconf

import (
	"runtime"
	"strconv"

	"github.com/asr-go/logrusconf/caller"
	"github.com/asr-go/logrusconf/hooks"
	"github.com/sirupsen/logrus"
)

// Init 初始化
func Init(level logrus.Level, filename string) {
	initFormatter(level)
	initHook(filename)
}

func initFormatter(level logrus.Level) {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.FullTimestamp = true                    // 显示完整时间
	customFormatter.TimestampFormat = "2006-01-02 15:04:05" // 时间格式
	customFormatter.DisableTimestamp = false                // 禁止显示时间
	customFormatter.DisableColors = false                   // 禁止颜色显示
	customFormatter.CallerPrettyfier = func(f *runtime.Frame) (string, string) {
		file, line := caller.GetCallerIgnoringLogMulti(2)
		return strconv.Itoa(line), file
	}

	logrus.SetReportCaller(true)
	logrus.SetFormatter(customFormatter)
	logrus.SetLevel(level)
}

func initHook(filename string) {
	logrus.AddHook(hooks.NewFileHook(filename))
}
