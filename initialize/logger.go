package initialize

import (
	"fmt"
	"github.com/develop-kevin/easy-gin-vue-admin/global"
	"gorm.io/gorm/logger"
)

type Writer struct {
	logger.Writer
}

// NewWriter writer 构造函数
func NewWriter(w logger.Writer) *Writer {
	return &Writer{Writer: w}
}

// Printf 格式化打印日志
func (w *Writer) Printf(message string, data ...interface{}) {
	var logZap bool
	logZap = global.EGVA_CONFIG.GormDB.LogZap
	if logZap {
		global.EGVA_LOG.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
