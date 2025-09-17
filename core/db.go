package core

import (
	"fmt"
	"rbacAdmin/global"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitGorm() (database *gorm.DB) {
	var db = global.Config.DB
	var dialector gorm.Dialector

	// 根据数据库类型构建连接字符串
	switch db.MODE {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			db.USER, db.PASSWORD, db.HOST, db.PORT, db.DbNAME)
		dialector = mysql.Open(dsn)
	case "pgsql", "postgres", "postgresql":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
			db.HOST, db.USER, db.PASSWORD, db.DbNAME, db.PORT)
		dialector = postgres.Open(dsn)
	case "sqlite":
		dialector = sqlite.Open(db.HOST)
	case "":
		logrus.Fatalf("未配置数据库类型")
		return nil

	default:
		logrus.Fatalf("不支持的数据库类型: %s", db.MODE)
	}

	// 打开数据库连接
	database, err := gorm.Open(dialector, &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 不生成外键约束
	})
	if err != nil {
		logrus.Fatalf("❌ 数据库连接失败: %v", err)
		return nil
	}

	// 配置连接池
	sqlDB, err := database.DB()
	if err != nil {
		logrus.Fatalf("获取数据库连接池失败: %v", err)
		return nil
	}

	// 测试数据库连接
	err = sqlDB.Ping()
	if err != nil {
		logrus.Fatalf("❌ 数据库连接测试失败: %v", err)
		return nil
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)           // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)          // 最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接最大生命周期

	logrus.Infof("✅ 数据库连接成功: %s", db.MODE)
	return database
}
