package main

import (
	"github.com/KarinaKuhne/API-usuario/database"
	"github.com/KarinaKuhne/API-usuario/routes"
)

func main() {
	database.ConnectDB()
	routes.HandleRequest()
}
