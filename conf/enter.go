package conf

type Config struct {
	Pgsql  Pgsql  `yaml:"pgsql"`
	Logger Logger `yaml:"logger"`
	System System `yaml:"system"`
}
