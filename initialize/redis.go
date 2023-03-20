package initialize

import (
	"context"
	"fmt"
	"github.com/develop-kevin/easy-gin-vue-admin/global"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func InitRedis() {
	rds := global.EGVA_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", rds.Host, rds.Port),
		Password: rds.Password,
		DB:       rds.DataBase,
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.EGVA_LOG.Error("redis connect ping failed,error:", zap.Error(err))
	} else {
		global.EGVA_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		pool := goredis.NewPool(client)
		global.EGVA_REDISCLIENT = client
		global.EGVA_REDISSYNC = redsync.New(pool)
	}
}
