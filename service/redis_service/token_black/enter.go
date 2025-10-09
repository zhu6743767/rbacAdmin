package token_black

import (
	"context"
	"rbacAdmin/global"
)

func HaveToken(token string) bool {
	// 从Redis中获取token
	_, err := global.Redis.Get(context.Background(), "black_"+token).Result()
	if err != nil {
		return false
	}
	return true
}
