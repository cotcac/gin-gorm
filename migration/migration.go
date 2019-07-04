package migration

import (
	"../models"
	"fmt"
)

func Migrate() {
	db:= models.DBConn()
	db.AutoMigrate(&models.User{},&models.Article{})
	fmt.Print("im migrate!!")
}