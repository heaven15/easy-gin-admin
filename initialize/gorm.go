package initialize

import (
	"github.com/develop-kevin/easy-gin-vue-admin/global"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var GormDBNew = new(gormDB)

type gormDB struct{}

// Config gorm 自定义配置
func (g *gormDB) Config(prefix string, singular bool) *gorm.Config {
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix,
			SingularTable: singular,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	loggerDefault := logger.New(NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})
	switch GetLogMode() {
	case "silent", "Silent":
		gormConfig.Logger = loggerDefault.LogMode(logger.Silent)
	case "error", "Error":
		gormConfig.Logger = loggerDefault.LogMode(logger.Error)
	case "warn", "Warn":
		gormConfig.Logger = loggerDefault.LogMode(logger.Warn)
	case "info", "Info":
		gormConfig.Logger = loggerDefault.LogMode(logger.Info)
	default:
		gormConfig.Logger = loggerDefault.LogMode(logger.Info)
	}
	return gormConfig
}

func GetLogMode() string {
	return global.EGVA_CONFIG.GormDB.LogMode
}
