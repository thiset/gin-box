package conf

import "strconv"

type Pgsql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DB       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log-Level"` // 日志等级
}

func (p Pgsql) Dsn() string {
	return "host=" + p.Host + " user=" + p.User + " password=" + p.Password +
		" dbname=" + p.DB + " port=" + strconv.Itoa(p.Port) +
		" sslmode=disable TimeZone=Asia/Shanghai"
}
