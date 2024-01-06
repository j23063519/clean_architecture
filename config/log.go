package config

type Log struct {
	TYPE      string `mapstructure:"LOG_TYPE"`       // single \ daily
	LEVEL     string `mapstructure:"LOG_LEVEL"`      // info \ error
	FILENAME  string `mapstructure:"LOG_FILENAME"`   // storage/logs/logs.log
	MAXSIZE   int    `mapstructure:"LOG_MAX_SIZE"`   // Unit: MB
	MAXBACKUP int    `mapstructure:"LOG_MAX_BACKUP"` // if 0 then not limited but if expire then still delete the log file
	MAXAGE    int    `mapstructure:"LOG_MAX_AGE"`    // if 7 then we will delete the log file after the 7 days and if 0 then we never delete the log file
	COMPRESS  bool   `mapstructure:"LOG_COMPRESS"`   // default: false, means not compressed because I need to read the log file
}
