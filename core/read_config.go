package core

import (
	"os"
	"rbacAdmin/config"
	"rbacAdmin/flags"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

func ReadConfig() *config.Config {
	byteData, err := os.ReadFile(flags.FlagOptions.File)
	if err != nil {
		logrus.Fatalf("❌ 配置文件读取失败: %v", err.Error())
		return nil
	}
	var c *config.Config
	err = yaml.Unmarshal(byteData, &c)
	if err != nil {
		logrus.Fatalf("❌ 配置文件格式解析失败: %v", err.Error())
		return nil
	}
	logrus.Infof("✅ 配置文件加载成功: %s", flags.FlagOptions.File)
	return c
}
