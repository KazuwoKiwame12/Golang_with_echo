package db

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Connect() *gorm.DB {
	//connstr := "host=" + os.Getenv("DB_HOST") + " port=" + os.Getenv("DB_PORT") + " user=" + os.Getenv("DB_USER") + " dbname=" + os.Getenv("DB_NAME") + " password=" + os.Getenv("DB_PASSWORD")
	connstr := os.Getenv("DATABASE_URL")
	db, err := gorm.Open("postgres", connstr)
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
