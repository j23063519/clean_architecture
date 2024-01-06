package config

type Redis struct {
	HOST     string `mapstructure:"REDIS_HOST"`
	PORT     string `mapstructure:"REDIS_PORT"`
	PASSWORD string `mapstructure:"REDIS_PASSWORD"`
	MAINDB   int    `mapstructure:"REDIS_MAIN_DB"`
	CACHEDB  int    `mapstructure:"REDIS_CACHE_DB"`
}
