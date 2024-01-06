package config

type App struct {
	NAME     string `mapstructure:"APP_NAME"`
	ENV      string `mapstructure:"APP_ENV"`
	KEY      string `mapstructure:"APP_KEY"`
	HTTP     string `mapstructure:"APP_HTTP"`
	HOST     string `mapstructure:"APP_HOST"`
	PORT     string `mapstructure:"APP_PORT"`
	URL      string `mapstructure:"APP_URL"`
	BASEPATH string `mapstructure:"APP_BASE_PATH"`
	TIMEZONE string `mapstructure:"APP_TIMEZONE"`
}
