package core

import (
	"fmt"
	"os"
	"rbacAdmin/config"

	"gopkg.in/yaml.v3"
)

func SetConfig(c *config.Config) {
	byteData, _ := yaml.Marshal(c)
	err := os.WriteFile("settings.yaml", byteData, 0666)
	if err != nil {
		fmt.Println("❌ 配置文件写入失败:", err)
		return
	}
	fmt.Println("✅ 配置保存成功:", c)
}
