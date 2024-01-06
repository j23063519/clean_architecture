package cmd

import (
	"time"

	"github.com/j23063519/clean_architecture/pkg/cache"
	"github.com/spf13/cobra"
)

var CmdPlay = &cobra.Command{
	Use:   "play",
	Short: "testing",
	Run:   runPlay,
	Args:  cobra.NoArgs,
}

// remember to remove the content of runPlay
func runPlay(cmd *cobra.Command, args []string) {
	// this is example
	cache.Set("test", "play", time.Duration(1)*time.Minute)
}
