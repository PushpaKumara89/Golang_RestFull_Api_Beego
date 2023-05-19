package main

import (
	"ApiBeeGo/db"
	_ "ApiBeeGo/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	db.DBConnection()
	beego.Run()
}
