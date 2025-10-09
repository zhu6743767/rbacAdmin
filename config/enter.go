package config

type Config struct {
	System  SystemConfig `yaml:"system"`
	DB      DB           `yaml:"db"`
	Redis   Redis        `yaml:"redis"`
	JWT     JWT          `yaml:"jwt"`
	Captcha Captcha      `yaml:"captcha"`
	Email   Email        `yaml:"email"`
	Upload  UploadConfig `yaml:"upload"`
}
