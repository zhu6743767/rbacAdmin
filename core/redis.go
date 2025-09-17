package core

import (
	"context"
	"rbacAdmin/global"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

func InitRedis() *redis.Client {
	r := global.Config.Redis
	c := redis.NewClient(&redis.Options{
		Addr:     r.Addr,
		Password: r.Password,
		DB:       r.DB,
	})
	if _, err := c.Ping(context.Background()).Result(); err != nil {
		logrus.Fatalf("❌ Redis连接失败: %v", err)
		return nil
	}
	logrus.Infof("✅ Redis连接成功: %s", r.Addr)
	return c
}
