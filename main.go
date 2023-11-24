package main

import (
	"github.com/hacktiv8-fp-golang/final-project-03/internal/database"
	"github.com/hacktiv8-fp-golang/final-project-03/internal/router"
)

// @title           Kanban Board
// @version         1.0
// @description     API for managing projects, tasks, and users within a project management system.
// @host            final-project-03-production.up.railway.app/

func main() {
	database.StartDB()
	router.StartServer()
}
