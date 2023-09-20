package main

import (
	"github.com/Godofin/travel-planner/database"
	"github.com/Godofin/travel-planner/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
