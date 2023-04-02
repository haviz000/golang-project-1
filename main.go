package main

import (
	"github.com/gin-gonic/gin"
	"github.com/haviz000/golang-project-1/database"
	"github.com/haviz000/golang-project-1/routes"
)

const PORT = ":8080"

func main() {
	router := gin.Default()

	database.ConnectDB()
	db := database.GetDB()

	routes.BookRoute(router, db)

	router.Run(PORT)
}
