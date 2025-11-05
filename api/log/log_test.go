package log

import (
	"compus-second-hand/global"
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	file, _ := os.OpenFile("./log.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	LogerInit(file)
	global.Logger.Error("hello world", "error", "test")
}
