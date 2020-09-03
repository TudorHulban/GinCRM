package cachesessionid

import (
	"os"

	"github.com/TudorHulban/GinCRM/pkg/cache"
	"github.com/TudorHulban/GinCRM/pkg/ostop"
	"github.com/TudorHulban/badgerwrap"
	"github.com/TudorHulban/log"
)

var theCache *badgerwrap.BStore

// GetCacheSessionsID Returns session ID cache object.
func GetCacheSessionsID() cache.IKV {
	if theCache == nil {
		l := log.New(log.DEBUG, os.Stderr, true)
		var errCo error
		theCache, errCo = badgerwrap.NewBStoreInMem(l)
		l.Info("Could not create Session ID Cache: ", errCo)
		os.Exit(ostop.CACHE)
	}
	return theCache
}
