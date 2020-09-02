package sessioncachebadger

import "github.com/TudorHulban/log"

// CacheBadger Provides struct for implementing model to satisfy session ID cache interface.
type CacheBadger struct {
	logger *log.LogInfo
}

// NewBadgerSessionIDCache Constructor for session ID cache.
func NewBadgerSessionIDCache() *CacheBadger {
	return nil
}
