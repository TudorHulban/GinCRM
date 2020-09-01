package sessioncache

// ISessionCache is the interface proposed for interacting with a session ID cache.
type ISessionCache interface {
	InsertSession(sessionID int64)                       // InsertSession inserts session ID
	InsertSessionWithTTL(sessionID int64, ttlSecs int64) // InsertSession inserts session ID with time to live in seconds
	FoundSessionID(sessionID int64) bool                 // FoundSessionID returns true if passed session ID is found
	DeleteSession(sessionID int64)                       // DeleteSession deletes passed session ID
}
