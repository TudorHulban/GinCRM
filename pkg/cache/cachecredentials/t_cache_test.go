package cachecredentials_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/TudorHulban/GinCRM/pkg/cache/cachecredentials"
	"github.com/TudorHulban/log"
)

func TestCacheWithLog(t *testing.T) {
	c, errCache := cachecredentials.GetCache(log.New(log.DEBUG, os.Stderr, true))
	require.Nil(t, errCache)
	assert.Nil(t, c.Close())
}

func TestCacheNoLog(t *testing.T) {
	_, errCache := cachecredentials.GetCache()
	require.Error(t, errCache)
}
