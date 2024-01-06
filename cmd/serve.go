package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/j23063519/clean_architecture/config"
	initail "github.com/j23063519/clean_architecture/init"
	"github.com/j23063519/clean_architecture/pkg/console"
	"github.com/spf13/cobra"
)

var CmdServe = &cobra.Command{
	Use:   "serve",
	Short: "start web server",
	Run:   runServe,
}

func runServe(cmd *cobra.Command, args []string) {
	gin.SetMode(gin.ReleaseMode)

	// new gin
	router := gin.New()

	// register route
	initail.SetRoute(router)

	// run server
	if err := router.Run(":" + config.Config.App.PORT); err != nil {
		console.Error("web server error: " + err.Error())
	}
}
