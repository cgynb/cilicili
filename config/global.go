package config

type Config struct {
	PageSize    int         `toml:"page_size"`
	TokenConfig tokenConfig `toml:"token"`
	MysqlConfig mysqlConfig `toml:"mysql"`
	RedisConfig redisConfig `toml:"redis"`
}
type tokenConfig struct {
	SecretKey  string `toml:"secret_key"`
	EffectTime int64  `toml:"effect_time"`
}

type mysqlConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DNS      string
}
type redisConfig struct {
	Host string
	Port string
	DB   int
}

var Conf Config
