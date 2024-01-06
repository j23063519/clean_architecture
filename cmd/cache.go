package cmd

import (
	"fmt"

	"github.com/j23063519/clean_architecture/pkg/cache"
	"github.com/j23063519/clean_architecture/pkg/console"
	"github.com/spf13/cobra"
)

var CmdCache = &cobra.Command{
	Use:   "cache",
	Short: "Cache management",
}

var CmdCacheClear = &cobra.Command{
	Use:   "clear",
	Short: "Clear cache",
	Run:   runCacheClear,
}

var CmdCacheForget = &cobra.Command{
	Use:   "forget",
	Short: "Delete redis key, example: cache forget cache-key",
	Run:   runCacheForget,
}

// forget flag
var cachekey string

func init() {
	// registe cache subflag
	CmdCache.AddCommand(CmdCacheClear, CmdCacheForget)

	// set cache forget flag
	CmdCacheForget.Flags().StringVarP(&cachekey, "key", "k", "", "KEY of the cache")
	CmdCacheForget.MarkFlagRequired("key")
}

func runCacheClear(cmd *cobra.Command, args []string) {
	cache.Flush()
	console.Success("Cache cleared.")
}

func runCacheForget(cmd *cobra.Command, args []string) {
	cache.Forget(cachekey)
	console.Success(fmt.Sprintf("Cache key [%s] deleted.", cachekey))
}
