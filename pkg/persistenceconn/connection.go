package persistenceconn

import (
	"log"
	"os"

	"github.com/TudorHulban/GinCRM/pkg/ostop"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var theRDBMSConn *gorm.DB

// GetRDBMSConn Returns RDBMS connection handler.
func GetRDBMSConn() *gorm.DB {
	if theRDBMSConn == nil {
		var errCo error
		theRDBMSConn, errCo = gorm.Open(sqlite.Open("/home/tudi/ram/gorm.db"), &gorm.Config{
			DisableAutomaticPing: true,
		})

		if errCo != nil {
			log.Println("Could not create RDBMS Connection: ", errCo)
			os.Exit(ostop.RDBMSConnection)
		}
	}
	return theRDBMSConn
}
