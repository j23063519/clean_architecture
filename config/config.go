package config

// environment setting
var Config struct {
	App     App     `mapstructure:",squash"`
	DB      DB      `mapstructure:",squash"`
	Log     Log     `mapstructure:",squash"`
	Redis   Redis   `mapstructure:",squash"`
	JWT     JWT     `mapstructure:",squash"`
	Mail    Mail    `mapstructure:",squash"`
	Captcha Captcha `mapstructure:",squash"`
}
