package cgorm_test

import (
	"os"
	"testing"

	"github.com/TudorHulban/GinCRM/pkg/persistence"
	"github.com/TudorHulban/GinCRM/pkg/persistence/cgorm"
	"github.com/TudorHulban/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddUser(t *testing.T) {
	require.Nil(t, cgorm.MigrateDBSchema())

	tt := []struct {
		testName    string
		data        persistence.UserAuth
		shouldError bool
	}{
		{testName: "validation issues", data: persistence.UserAuth{
			UserCode: "abcd",
		}, shouldError: true},
		{testName: "valid data", data: persistence.UserAuth{
			SecurityGroupID: 1,
			UserCode:        "abcd",
		}, shouldError: false},
	}

	userCRUD := cgorm.NewUser(log.New(log.DEBUG, os.Stderr, true))

	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			if tc.shouldError {
				assert.Error(t, userCRUD.AddUser(&tc.data))
				return
			}
			assert.Nil(t, userCRUD.AddUser(&tc.data))
		})
	}
}
