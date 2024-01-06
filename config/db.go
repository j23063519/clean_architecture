package config

type DB struct {
	PGSQL PGSQL `mapstructure:",squash"`
}

type PGSQL struct {
	HOST     string `mapstructure:"DB_PGSQL_HOST"`
	PORT     int    `mapstructure:"DB_PGSQL_PORT"`
	DATABASE string `mapstructure:"DB_PGSQL_DATABASE"`
	USERNAME string `mapstructure:"DB_PGSQL_USERNAME"`
	PASSWORD string `mapstructure:"DB_PGSQL_PASSWORD"`
	SOURCE   string `mapstructure:"DB_PGSQL_SOURCE"`
	DEBUG    string `mapstructure:"DB_PGSQL_DEBUG"`
}
