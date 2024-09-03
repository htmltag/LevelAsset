package main

import (
	"assetdog/database"
	"assetdog/routes"

	"github.com/gin-gonic/gin"
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
