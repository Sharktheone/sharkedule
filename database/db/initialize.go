package db

import (
	"github.com/Sharktheone/sharkedule/database"
	jsonfileV2 "github.com/Sharktheone/sharkedule/database/jsonfile/V2"
	"log"
)

var (
	DB database.IDatabaseV2
)

func Init() {
	DB = jsonfileV2.NewJSONFile()
	if err := DB.Load(); err != nil {
		log.Fatalf("failed loading database v2: %v", err)
	}
}
