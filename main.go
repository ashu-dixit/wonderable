package main

import (
	"wonderable/config"
	"wonderable/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	router := gin.Default()
	routes.AuthRoutes(router)
	routes.AdminRoutes(router)
	router.Run(":8080")
}
