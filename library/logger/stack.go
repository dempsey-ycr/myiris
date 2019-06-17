package logger

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
	"sync/atomic"
)

const (
	SWITCH = 1 // 是否打印堆栈
	SKIP   = 4 // 过滤前四层
)

// global error config object
var config = &errConfig{}

// error config
type errConfig struct {
	isPrintStack uint32
}

//set error config info
func SetConfig(conf byte) {
	if (conf & SWITCH) != 0 {
		atomic.CompareAndSwapUint32(&config.isPrintStack, 0, 1)
	}
}

func StackTrace() string {
	return genStackTrace(SKIP)
}

func genStackTrace(skip int) string {
	if config.isPrintStack == 1 {
		var buffer bytes.Buffer
		var st [64]uintptr
		n := runtime.Callers(skip, st[:])
		callers := st[:n]
		frames := runtime.CallersFrames(callers)
		for {
			frame, ok := frames.Next()
			if !ok {
				break
			}
			if !strings.Contains(frame.File, "runtime/") {
				buffer.WriteString(fmt.Sprintf("%s\t%s:%d\n",
					frame.Func.Name(), frame.File, frame.Line))
			}
		}
		return buffer.String()
	}
	return ""
}
