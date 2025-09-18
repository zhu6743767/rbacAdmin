package config

type JWT struct {
	Secret string `yaml:"secret"`
	Expire int    `yaml:"expire"` // 单位小时
	Issuer string `yaml:"issuer"`
}
