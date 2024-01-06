package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// env file name，EX: .env、.env.prod、.env.dev
var Env string

// register a global flag in root cmd
func RegisterGlobalFlags(rootCmd *cobra.Command) {
	// rootCmd.PersistentFlags().StringVarP:
	// binding a global string type of Env with --env flag
	// we can use -e or --env as flag
	// last parameter is descrition: load .env file, example: --env=testing use .env.testing file
	rootCmd.PersistentFlags().StringVarP(&Env, "env", "e", "", "load .env file, example: --env=testing use .env.testing file")
}

// register a default command line
//
// rootCmd: root command line，subCmd: register a default command line
func RegisterDefaultCmd(rootCmd *cobra.Command, subCmd *cobra.Command) {
	cmd, _, err := rootCmd.Find(os.Args[1:])

	firstArg := ""
	if len(os.Args[1:]) > 0 {
		firstArg = os.Args[1]
	}

	if err != nil && cmd.Use == rootCmd.Use && firstArg != "-h" && firstArg != "--help" {
		args := append([]string{subCmd.Use}, os.Args[1:]...)

		rootCmd.SetArgs(args)
	}
}
