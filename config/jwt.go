package config

type JWT struct {
	EXPIRETIME     string `mapstructure:"JWT_EXPIRE_TIME"`
	MAXREFRESHTIME string `mapstructure:"JWT_MAX_REFRESH_TIME"`
}
