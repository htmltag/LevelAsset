package routes

import (
	"levelasset/controllers"

	"github.com/gin-gonic/gin"
	"github.com/syndtr/goleveldb/leveldb"
)

func AssetRoute(router *gin.Engine, db *leveldb.DB) {
	router.GET("/api/assets", controllers.GetAssets(db))
	router.POST("/api/assets", controllers.CreateAsset(db))
}
