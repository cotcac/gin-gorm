package main

import (
	"./router"
	"fmt"
	"./migration"
)

func main() {
	fmt.Print("run me first!")
	migration.Migrate()
	r:= router.Router()
	r.Run() // listen and serve on 0.0.0.0:8080
}