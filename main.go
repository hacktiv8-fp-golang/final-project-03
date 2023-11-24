package main

import (
	"github.com/hacktiv8-fp-golang/final-project-03/internal/database"
	"github.com/hacktiv8-fp-golang/final-project-03/internal/router"
)

// @title           Kanban Board
// @version         1.0
// @description     This is a simple server mimicking Instagram server.
// @host            localhost:8080

func main() {
	database.StartDB()
	router.StartServer()
}
