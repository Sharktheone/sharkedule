package db

import (
	"github.com/Sharktheone/sharkedule/database"
	"github.com/Sharktheone/sharkedule/database/jsonfile"
	jsonfileV2 "github.com/Sharktheone/sharkedule/database/jsonfile/V2"
	"log"
)

var (
	DB   database.IDatabase
	DBV2 database.IDatabaseV2
)

func Init() {
	DB = jsonfile.NewJSONFile()
	if err := DB.Load(); err != nil {
		log.Fatalf("failed loading database: %v", err)
	}
	DBV2 = jsonfileV2.NewJSONFile()
	if err := DBV2.Load(); err != nil {
		log.Fatalf("failed loading database v2: %v", err)
	}
}
