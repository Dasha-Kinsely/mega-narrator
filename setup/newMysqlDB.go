package setup

import (
	"github.com/Dasha-Kinsely/mega-narrator/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlDB(dsn string, disableDatetimePrecision bool, dontSupportRenameIndex bool, dontSupportRenameColumn bool, skipInitializeWithVersion bool) {
	// Database connection configuration but returns nothing
	db, dbErr := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  disableDatetimePrecision,
		DontSupportRenameIndex:    dontSupportRenameIndex,
		DontSupportRenameColumn:   dontSupportRenameColumn,
		SkipInitializeWithVersion: skipInitializeWithVersion,
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if dbErr != nil {
		// log.Println("Error initializing MYSQL_DB...")
		panic("Error initializing MYSQL_DB...")
	}
	models.DB = db
	MysqlMigration()
}

