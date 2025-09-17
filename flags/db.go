package flags

import (
	"rbacAdmin/global"
	"rbacAdmin/models"

	"github.com/sirupsen/logrus"
)

func AutoMigrate() {
	// 这里可以添加数据库自动迁移的逻辑
	err := global.DB.AutoMigrate(&models.UserModel{}, &models.RoleModel{}, &models.MenuModel{}, &models.APIModel{}, &models.RoleMenuModel{})
	if err != nil {
		logrus.Fatalf("数据库自动迁移失败: %v", err)
	}
	logrus.Info("数据库自动迁移成功")
}
