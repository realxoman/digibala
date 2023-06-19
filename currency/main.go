package main

import (
	"currency/db"
	"currency/server"
)

func main() {
	db.ConnectDatabase()
	server.StartServer()
}
