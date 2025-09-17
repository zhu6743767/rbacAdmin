package core

import (
	"rbacAdmin/global"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/sirupsen/logrus"
)

func InitCasbin() *casbin.CachedEnforcer {
	a, _ := gormadapter.NewAdapterByDB(global.DB)
	casbinModel := `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
`
	m, err := model.NewModelFromString(casbinModel)
	if err != nil {
		logrus.Fatalf("初始化Casbin失败: %v", err)
	}
	e, _ := casbin.NewCachedEnforcer(m, a)
	e.SetExpireTime(60 * 60)
	_ = e.LoadPolicy()
	return e
}
