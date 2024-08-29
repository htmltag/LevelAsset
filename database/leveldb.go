package database

import (
	"fmt"
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

func ConnectDB() *leveldb.DB {
	db, err := leveldb.OpenFile("mydb", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to LevelDB!")

	return db
}

// DB instance
var DB *leveldb.DB = ConnectDB()
