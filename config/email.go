package config

type Email struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
}

func (e Email) Verify() bool {
	if e.User == "" || e.Password == "" || e.Host == "" || e.Port == 0 {
		return false
	}
	return true
}
