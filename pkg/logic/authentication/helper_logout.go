package authentication

import "github.com/TudorHulban/GinCRM/pkg/cache/cachelogin"

// logoutCredentialsCache Helper deletes credentials to clean up user access.
func logoutCredentialsCache(usercode string) error {
	return cachelogin.GetCache().DeleteKVByK([]byte(usercode))
}
