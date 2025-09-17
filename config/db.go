package config

type DB struct {
	MODE     string `yaml:"mode"`
	HOST     string `yaml:"host"`
	PORT     int    `yaml:"port"`
	USER     string `yaml:"user"`
	PASSWORD string `yaml:"password"`
	DbNAME   string `yaml:"db_name"`
}
