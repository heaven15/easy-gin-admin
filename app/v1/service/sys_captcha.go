package service

import (
	"context"
	"github.com/develop-kevin/easy-gin-vue-admin/global"
	"github.com/mojocn/base64Captcha"
	"time"
)

func NewDefaultRedisStore() *redisStore {
	return &redisStore{
		Expiration: time.Hour * 24,
		PreKey:     "SysCaptCha",
	}
}

type redisStore struct {
	Expiration time.Duration
	PreKey     string
	Context    context.Context
}

func (rs *redisStore) UseWithCtx(ctx context.Context) base64Captcha.Store {
	rs.Context = ctx
	return rs
}

func (rs *redisStore) Set(id, value string) error {
	if err := global.EGVA_REDISCLIENT.Set(rs.Context, rs.PreKey+id, value, rs.Expiration).Err(); err != nil {
		global.EGVA_LOG.Error("redis set captcha failed:" + err.Error())
		return err
	}
	return nil
}

func (rs *redisStore) Get(id string, clear bool) string {
	v, err := global.EGVA_REDISCLIENT.Get(rs.Context, id).Result()
	if err != nil {
		global.EGVA_LOG.Error("redis get captcha failed:" + err.Error())
		return ""
	}
	if clear {
		if err = global.EGVA_REDISCLIENT.Del(rs.Context, id).Err(); err != nil {
			global.EGVA_LOG.Error("redis del captcha failed:" + err.Error())
			return ""
		}
	}
	return v
}

func (rs *redisStore) Verify(id, answer string, clear bool) bool {
	key := rs.PreKey + id
	v := rs.Get(key, clear)
	if v == "" {
		return false
	}
	return v == answer
}
