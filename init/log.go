package init

import (
	"github.com/j23063519/clean_architecture/config"
	"github.com/j23063519/clean_architecture/pkg/log"
)

func SetLog() {
	log.NewLogger(
		config.Config.Log.FILENAME,
		config.Config.Log.MAXSIZE,
		config.Config.Log.MAXBACKUP,
		config.Config.Log.MAXAGE,
		config.Config.Log.COMPRESS,
		config.Config.Log.TYPE,
		config.Config.Log.LEVEL,
	)
}
