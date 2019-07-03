
package models 
import (
	"github.com/jinzhu/gorm"
	"time"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DBConn() (db *gorm.DB) {
	dbDriver := "mysql"
    dbUser := "root"
    dbPass := ""
    dbName := "gorm"
    db, err := gorm.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        panic(err.Error())
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	db.DB().SetConnMaxLifetime(time.Hour)
    // defer db.Close()
    return db
}