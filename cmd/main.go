package main

import (
	"fmt"

	"github.com/TalesPalma/api-go-rest/database"
	"github.com/TalesPalma/api-go-rest/routes"
)

func main() {
	database.ConectDatabase()
	fmt.Println("Iniciando servidor go")
	routes.HandleRequest()
}
