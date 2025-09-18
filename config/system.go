package config

import (
	"fmt"
)

type SystemConfig struct {
	Mode string `yaml:"mode"`
	Ip   string `yaml:"ip"`
	Port int    `yaml:"port"`
}

func (s *SystemConfig) Addr() string {
	return fmt.Sprintf("%s:%d", s.Ip, s.Port)
}
