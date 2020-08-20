package connect

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func connectDB() *gorm.DB {
	connstr := "host=" + os.Getenv("DB_HOST") + " port=" + os.Getenv("DB_PORT") + " user=" + os.Getenv("DB_USER") + " dbname=" + os.Getenv("DB_NAME") + " password=" + os.Getenv("DB_PASSWORD")
	db, err := gorm.Open("postgres", connstr)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	return db
}
