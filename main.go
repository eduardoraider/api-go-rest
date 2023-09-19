package main

import (
	"api-go-rest/database"
	"api-go-rest/routes"
	"fmt"
)

func main() {
	database.ConnectDatabase()
	fmt.Println("Iniciando servidor rest")
	routes.HandleRequest()
}
