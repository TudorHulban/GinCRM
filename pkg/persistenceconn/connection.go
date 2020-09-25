package persistenceconn

import (
	"log"
	"os"
	"strings"

	"github.com/TudorHulban/GinCRM/cmd/setup"
	"github.com/TudorHulban/GinCRM/pkg/logic/authentication"
	"github.com/TudorHulban/GinCRM/pkg/ostop"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var theRDBMSConn *gorm.DB

// GetRDBMSConn Returns RDBMS connection handler.
func GetRDBMSConn() *gorm.DB {
	if theRDBMSConn == nil {
		var errCo error

		os.MkdirAll(setup.SQLiteFilePath, os.ModePerm)
		path := strings.Split(setup.SQLiteFilePath, ".")
		theRDBMSConn, errCo = gorm.Open(sqlite.Open(strings.Join([]string{path[0], authentication.UXNano(), ".", path[1]}, "")), &gorm.Config{
			DisableAutomaticPing: true,
		})

		if errCo != nil {
			log.Println("Could not create RDBMS Connection: ", errCo)
			os.Exit(ostop.RDBMSConnection)
		}
	}
	return theRDBMSConn
}
