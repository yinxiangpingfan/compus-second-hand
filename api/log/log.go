package log

import (
	"compus-second-hand/global"
	"log/slog"
	"os"
)

//用于记录日志

func LogerInit(file *os.File) {
	global.Logger = slog.New(slog.NewJSONHandler(file, nil))
	slog.SetDefault(
		global.Logger,
	)
}
