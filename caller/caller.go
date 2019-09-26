package caller

import (
	"runtime"
	"strings"
)

// getCaller returns the filename and the line info of a function
// further down in the call stack.  Passing 0 in as callDepth would
// return info on the function calling getCallerIgnoringLog, 1 the
// parent function, and so on.  Any suffixes passed to getCaller are
// path fragments like "/pkg/log/log.go", and functions in the call
// stack from that file are ignored.
func getCaller(callDepth int, ignoresArray ...[]string) (file string, line int) {
	callDepth++
outer:
	for {
		var ok bool
		_, file, line, ok = runtime.Caller(callDepth)
		if !ok {
			file = "???"
			line = 0
			break
		}

		for _, ignores := range ignoresArray {
			include := true
			for _, s := range ignores {
				if strings.Index(file, s) == -1 {
					include = false
				}
			}

			if include {
				callDepth++
				continue outer
			}
		}
		break
	}
	array := strings.Split(file, "/")
	file = array[len(array)-1]
	return
}

// GetCallerIgnoringLogMulti GetCallerIgnoringLogMulti
func GetCallerIgnoringLogMulti(callDepth int) (string, int) {
	return getCaller(callDepth+1, []string{"logrus", "hooks.go"}, []string{"logrus", "entry.go"}, []string{"logrus", "logger.go"}, []string{"logrus", "exported.go"}, []string{"proc.go"}, []string{"asm_amd64.s"})
}
