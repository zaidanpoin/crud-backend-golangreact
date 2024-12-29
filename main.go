package main

import (
	model "github.com/zaidanpoin/crud-golang-react/Model"
	routes "github.com/zaidanpoin/crud-golang-react/Routes"
	"github.com/zaidanpoin/crud-golang-react/database"
)

func main() {
	loadDatabase()
	routes.ServeApps()

}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.Member{})

}
