package flags

import (
	"flag"
	"os"
)

type Options struct {
	DB   bool
	File string
	Menu string
	Type string
}

var FlagOptions Options

func init() {
	flag.StringVar(&FlagOptions.File, "f", "settings.yaml", "配置文件路径")
	flag.StringVar(&FlagOptions.Menu, "m", "menu", "菜单")
	flag.StringVar(&FlagOptions.Type, "t", "", "类型")
	flag.BoolVar(&FlagOptions.DB, "db", false, "数据库迁移")
	flag.Parse()
}

func Run() {
	if FlagOptions.DB {
		AutoMigrate()
		os.Exit(0)
	}
}
