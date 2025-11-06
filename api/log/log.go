package log

import (
	"compus-second-hand/global"
	"log/slog"
	"os"
)

//用于记录日志

func logerInit(file *os.File) {
	global.Logger = slog.New(slog.NewJSONHandler(file, nil))
	slog.SetDefault(
		global.Logger,
	)
}

//打开日志文件并初始化日志

func InitLog(path string) *os.File {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("log日志初始化失败: " + err.Error())
	}
	logerInit(f)
	return f
}
