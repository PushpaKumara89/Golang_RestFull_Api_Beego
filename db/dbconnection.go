package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() *sql.DB {
	//https://stackoverflow.com/questions/36256230/connection-fails-with-mysql-using-golang

	// Open up our database connection.
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3307)/GoTextDb")

	// if there is an error opening the connection, handle it
	if err != nil {
		println("Connection false")
		log.Print(err.Error())
	} else {
		log.Print("Connected--> mysql in localhost:3307")
		// println("Connected")
	}
	return db
}
