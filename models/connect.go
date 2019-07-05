
package models 
import (
	"github.com/jinzhu/gorm"
	"time"
	"encoding/json"
	"os"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
type Config struct {
	Database struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Db string `json:"db"`
	}`json:"database"`
}
func LoadConfiguration(filename string) (Config, error) {
	var config Config
	configFile, err := os.Open(filename)
	defer configFile.Close()
    if err!=nil{
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err

}
func DBConn() (db *gorm.DB) {
	config,err := LoadConfiguration("config.json")
	if err != nil {
		panic(err.Error())
	}
	dbDriver := "mysql"
    dbUser := config.Database.Username
    dbPass := config.Database.Password
    dbName := config.Database.Db
    db, err = gorm.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
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