package config

type UploadConfig struct {
	Size int64  `yaml:"size"`
	Ext  string `yaml:"ext"`
	Dir  string `yaml:"dir"`
}
