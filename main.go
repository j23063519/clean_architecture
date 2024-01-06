package main

import (
	"fmt"
	"os"

	"github.com/j23063519/clean_architecture/cmd"
	initail "github.com/j23063519/clean_architecture/init"
	cmdpkg "github.com/j23063519/clean_architecture/pkg/cmd"
	configpkg "github.com/j23063519/clean_architecture/pkg/config"
	"github.com/j23063519/clean_architecture/pkg/console"
	"github.com/spf13/cobra"
)

//	@termsOfService	http://swagger.io/terms/
//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey	Bearer
// @in							header
// @name						Authorization
// @description				Type "Bearer" followed by a space and JWT token.
func main() {
	var rootCmd = &cobra.Command{
		Use:   "gocla",
		Short: "golang clean architecture example",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,
		// rootCmd : execute all subcommands
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// init setting --env parameter
			configpkg.InitConfig(".", cmdpkg.Env)

			// init logger
			initail.SetLog()

			// init db
			initail.SetDB()

			// init redis
			initail.SetRedis()

			// init cache
			initail.SetCache()
		},
	}

	// register subcommand
	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdMigrate,
		cmd.CmdDBSeed,
		cmd.CmdKey,
		cmd.CmdPlay,
		cmd.CmdCache,
	)

	// set default default command (excution server)
	cmdpkg.RegisterDefaultCmd(rootCmd, cmd.CmdServe)

	// register a global environment --env
	cmdpkg.RegisterGlobalFlags(rootCmd)

	// excute root command
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("failed to run app with %v: %s", os.Args, err.Error()))
	}
}
