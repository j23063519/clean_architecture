package config

import (
	"log/slog"
	"os"

	conf "github.com/j23063519/clean_architecture/config"

	viperlib "github.com/spf13/viper"
)

// viper instance
var viper *viperlib.Viper

// InitConfig loads the configuration from a config file or environment variables.
func InitConfig(path, envSuffix string) {
	// init viper
	viper = viperlib.New()

	// "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
	viper.SetConfigType("env")

	viper.AddConfigPath(path)

	// read env file
	viper.AutomaticEnv()

	setConfigName(envSuffix)

	err := viper.ReadInConfig()
	if err != nil {
		slog.Error("Could not read config: %v", err)
		return
	}

	// read env and then binding data to Config
	err = viper.Unmarshal(&conf.Config)
	if err != nil {
		slog.Error("Could not read config: %v", err)
		return
	}
}

// set and watch environment files
func setConfigName(envSuffix string) {
	// default .env file ，if --env=name，load .env.name file
	envPath := ".env"
	if len(envSuffix) > 0 {
		filePath := ".env." + envSuffix

		// if find file，then envPath = filePath
		if _, err := os.Stat(filePath); err == nil {
			// 如 .env.testing or .env.stage
			envPath = filePath
		}
	}

	viper.SetConfigName(envPath)

	// watch env
	viper.WatchConfig()
}
