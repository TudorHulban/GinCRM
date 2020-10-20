package cachesessionid

import (
	"os"

	"github.com/TudorHulban/GinCRM/pkg/cache"
	"github.com/TudorHulban/GinCRM/pkg/ostop"
	"github.com/TudorHulban/badgerwrap"
	"github.com/TudorHulban/log"
)

var theCache *badgerwrap.BStore

// GetCache Returns session ID cache object.
// Using varidic for the cases where logger for sure was already created.
func GetCache(l ...*log.LogInfo) cache.IKV {
	if (theCache == nil) && (len(l) > 0) {
		var errCo error
		theCache, errCo = badgerwrap.NewBStoreInMem(l[0])
		l[0].Infof("error trying to create session ID cache:%v", errCo)
		os.Exit(ostop.CACHE)
	}
	return theCache
}
