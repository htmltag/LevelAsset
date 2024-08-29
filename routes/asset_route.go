package routes

import (
	"assetdog/controllers"

	"github.com/gin-gonic/gin"
	"github.com/syndtr/goleveldb/leveldb"
)

func AssetRoute(router *gin.Engine, db *leveldb.DB) {
	router.GET("/assets", controllers.GetAssets(db))
	router.POST("/assets", controllers.CreateAsset(db))
}
