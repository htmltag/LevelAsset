package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

		// if id is not provided, generate a new one
		if _, ok := j["id"]; !ok {
			j["id"] = uuid.New().String()
		}

		// if in need of required fields
		// if _, ok := j["name"]; !ok {
		// 	c.JSON(400, gin.H{
		// 		"message": "Name is required",
		// 	})
		// 	return
		// }

		err = db.Put([]byte(j["id"].(string)), []byte(jd), nil)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Error creating asset",
			})
			return
		}

		c.JSON(http.StatusCreated, j)
	}
}
