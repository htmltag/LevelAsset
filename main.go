package main

import (
	// Add required Go packages
	"github.com/gin-gonic/gin"

	// Add the MongoDB driver packages
	"assetdog/database"
	"assetdog/routes"
)

func main() {
	db := database.DB
	defer db.Close()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	routes.AssetRoute(r, db)

	r.Run()
}
