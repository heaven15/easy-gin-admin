package initialize

import (
	"fmt"
	"github.com/develop-kevin/easy-gin-vue-admin/global"
	"github.com/develop-kevin/easy-gin-vue-admin/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// InitZap 获取 zap.Logger
func InitZap() {
	if ok, _ := utils.PathExists(global.EGVA_CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.EGVA_CONFIG.Zap.Director)
		_ = os.Mkdir(global.EGVA_CONFIG.Zap.Director, os.ModePerm)
	}
	cores := utils.ZapNew.GetZapCores()
	global.EGVA_LOG = zap.New(zapcore.NewTee(cores...))
	if global.EGVA_CONFIG.Zap.ShowLine {
		global.EGVA_LOG = global.EGVA_LOG.WithOptions(zap.AddCaller())
	}
	zap.ReplaceGlobals(global.EGVA_LOG)
}
