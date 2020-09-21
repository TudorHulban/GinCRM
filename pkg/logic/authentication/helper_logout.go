package authentication

import "github.com/TudorHulban/GinCRM/pkg/cache/cachecredentials"

// logoutCredentialsCache Helper deletes credentials to clean up user access.
func logoutCredentialsCache(usercode string) error {
	return cachecredentials.GetCache().DeleteKVByK([]byte(usercode))
}
