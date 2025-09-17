package global

import (
	"rbacAdmin/config"

	"github.com/casbin/casbin/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	DB     *gorm.DB
	Redis  *redis.Client
	Casbin *casbin.CachedEnforcer
)
