package migration

import (
	"../models"
	"fmt"
)

func Migrate() {
	db:= models.DBConn()
	db.AutoMigrate(&models.Article{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Category{})
	fmt.Print("im migrate!!")
}