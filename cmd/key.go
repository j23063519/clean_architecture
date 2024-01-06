package cmd

import (
	"github.com/j23063519/clean_architecture/pkg/console"
	"github.com/j23063519/clean_architecture/pkg/util"
	"github.com/spf13/cobra"
)

var CmdKey = &cobra.Command{
	Use:   "key",
	Short: "Generate key",
	Run:   runGenerateKey,
	Args:  cobra.NoArgs,
}

func runGenerateKey(cmd *cobra.Command, args []string) {
	console.Error("---")
	console.Success("App key:" + util.RandomString(32))
	console.Success("---")
}
