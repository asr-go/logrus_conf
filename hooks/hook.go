package hooks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/asr-go/logrusconf/caller"
)

// FileHook FileHook
type FileHook struct {
	W *fileLogWriter
}

// fileHookConfig fileHookConfig
type fileHookConfig struct {
	Filename   string `json:"filename"`
	MaxLines   int64  `json:"maxLines"`
	MaxSize    int64  `json:"maxsize"`
	Daily      bool   `json:"daily"`
	MaxDays    int64  `json:"maxDays"`
	Rotate     bool   `json:"rotate"`
	Perm       string `json:"perm"`
	RotatePerm string `json:"rotateperm"`
	Level      int32  `json:"level"`
}

// NewFileHook NewFileHook
func NewFileHook(filename string) (hook logrus.Hook) {

	dir := filepath.Dir(filename)

	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return
	}

	hookConf := fileHookConfig{
		Filename:   filename,
		Daily:      true,
		MaxDays:    7,
		Rotate:     true,
		MaxLines:   10000,
		MaxSize:    2048000,
		RotatePerm: "0440",
		Perm:       "0660",
		Level:      3,
	}

	w := newFileWriter()

	confData, err := json.Marshal(hookConf)

	if err != nil {
		return
	}

	err = w.Init(string(confData))
	if err != nil {
		return
	}

	hook = &FileHook{W: w}

	return
}

// Fire Fire
func (p *FileHook) Fire(entry *logrus.Entry) (err error) {
	message, err := getMessage(entry)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read entry, %v", err)
		return err
	}

	now := time.Now()

	file, lineNumber := caller.GetCallerIgnoringLogMulti(2)
	position := fmt.Sprintf("%s%s%d", file, ":", lineNumber)
	position = fmt.Sprintf("%s%24v%s", "[", position, "]")
	message = fmt.Sprintf("%v ", position) + message

	switch entry.Level {
	case logrus.PanicLevel:
		return p.W.WriteMsg(now, fmt.Sprintf("[PANI] %s", message), LevelError)
	case logrus.FatalLevel:
		return p.W.WriteMsg(now, fmt.Sprintf("[FATA] %s", message), LevelError)
	case logrus.ErrorLevel:
		return p.W.WriteMsg(now, fmt.Sprintf("[ERRO] %s", message), LevelError)
	case logrus.WarnLevel:
		return p.W.WriteMsg(now, fmt.Sprintf("[WARN] %s", message), LevelWarn)
	case logrus.InfoLevel:
		return p.W.WriteMsg(now, fmt.Sprintf("[INFO] %s", message), LevelInfo)
	case logrus.DebugLevel:
		return p.W.WriteMsg(now, fmt.Sprintf("[DEBU] %s", message), LevelDebug)
	default:
		return nil
	}
}

// Levels Levels
func (p *FileHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}

// getMessage getMessage
func getMessage(entry *logrus.Entry) (message string, err error) {
	message = message + fmt.Sprintf("%s", entry.Message)

	for k, v := range entry.Data {
		if !strings.HasPrefix(k, "err_") {
			message = message + fmt.Sprintf(" %v=%v", k, v)
		}
	}

	if errCode, exist := entry.Data["err_code"]; exist {

		ns, _ := entry.Data["err_ns"]
		ctx, _ := entry.Data["err_ctx"]
		id, _ := entry.Data["err_id"]
		tSt, _ := entry.Data["err_stack"]
		st, _ := tSt.(string)
		st = strings.Replace(st, "\n", "\n\t\t", -1)

		buf := bytes.NewBuffer(nil)
		buf.WriteString(fmt.Sprintf("\tid:\n\t\t%s#%d:%s\n", ns, errCode, id))
		buf.WriteString(fmt.Sprintf("\tcontext:\n\t\t%s\n", ctx))
		buf.WriteString(fmt.Sprintf("\tstacktrace:\n\t\t%s", st))

		message = message + fmt.Sprintf("%v", buf.String())
	}

	return
}
