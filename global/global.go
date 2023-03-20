package global

import (
	"github.com/develop-kevin/easy-gin-vue-admin/config"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-redsync/redsync/v4"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	EGVA_ENV         = "EGVA_CONFIG"
	EGVA_TRANS       ut.Translator
	EGVA_LOG         *zap.Logger
	EGVA_CONFIG      *config.Config
	EGVA_DB          *gorm.DB
	EGVA_REDISCLIENT *redis.Client
	EGVA_REDISSYNC   *redsync.Redsync
)
