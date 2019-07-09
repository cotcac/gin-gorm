package main

import (
	"fmt"

	"./migration"
	"./router"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	fmt.Print("run me first!")
	migration.Migrate()
	r := router.Router()
	r.Run() // listen and serve on 0.0.0.0:8080
}
