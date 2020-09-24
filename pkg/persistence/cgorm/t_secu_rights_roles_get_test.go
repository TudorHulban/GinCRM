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

func TestGetSecurityRightsByProfile(t *testing.T) {
	setup.CleanRDBMS()
	require.Nil(t, cgorm.MigrateDBSchema())

	l := log.New(log.DEBUG, os.Stderr, true)
	assert.Nil(t, cgorm.PopulateSchemaSecurityRoles(l))

	secu := cgorm.NewSecurityRR(l)

	userRights, errOp := secu.GetSecurityRightsForProfile(0)
	assert.Nil(t, errOp)
	assert.Greater(t, len(userRights), 0)
}
