package cgorm_test

import (
	"os"
	"testing"

	"github.com/TudorHulban/GinCRM/cmd/setup"
	"github.com/TudorHulban/GinCRM/pkg/persistence/cgorm"
	"github.com/TudorHulban/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDataPopulationRR(t *testing.T) {
	setup.CleanRDBMS()
	require.Nil(t, cgorm.MigrateDBSchema())

	l := log.New(log.DEBUG, os.Stderr, true)
	assert.Nil(t, cgorm.PopulateSchemaSecurityRoles(l))
}
