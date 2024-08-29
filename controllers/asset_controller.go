package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syndtr/goleveldb/leveldb"
)

func GetAssets(db *leveldb.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		iter := db.NewIterator(nil, nil)
		var asset map[string]interface{}
		for iter.Next() {
			err := json.Unmarshal(iter.Value(), &asset)
			if err != nil {
				c.JSON(500, gin.H{
					"message": "Error reading asset",
				})
				return
			}
		}
		iter.Release()

		// Return the assets as JSON
		c.JSON(http.StatusOK, asset)
	}
}

func CreateAsset(db *leveldb.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		jd, err := c.GetRawData()
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Invalid request",
			})
			return
		}

		j := make(map[string]interface{})
		json.Unmarshal(jd, &j)

		if _, ok := j["source"]; !ok {
			c.JSON(400, gin.H{
				"message": "Source is required",
			})
			return
		}

		err = db.Put([]byte(j["source"].(string)), []byte(jd), nil)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Error creating asset",
			})
			return
		}

		c.JSON(http.StatusCreated, j)
	}
}
