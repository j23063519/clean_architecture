package config

type Mail struct {
	HOST     string `mapstructure:"MAIL_HOST"`
	PORT     string `mapstructure:"MAIL_PORT"`
	USERNAME string `mapstructure:"MAIL_USERNAME"`
	PASSWORD string `mapstructure:"MAIL_PASSWORD"`
	FROMNAME string `mapstructure:"MAIL_FROM_NAME"`
}
