package database

import (
	"fmt"
	"log"
	"time"

	"github.com/gabrielteiga/startup-rush/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBGormAdapter struct {
	DB    *gorm.DB
	Error error
}

func InitConnection() *DBGormAdapter {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		configs.DB_USER,
		configs.DB_PASSWORD,
		configs.DB_DOMAIN,
		configs.DB_PORT,
		configs.DB_DATABASE,
	)

	log.Println(dsn)
	for i := 0; i < 10; i++ {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Connected to the database")
			return &DBGormAdapter{
				DB:    db,
				Error: err,
			}
		}
		log.Println("Retrying to connect to the database...")
		time.Sleep(6 * time.Second)
	}
	log.Println("Failed to connect to the database after 10 attempts")
	return nil
}

func (db *DBGormAdapter) Migrate() error {
	return db.DB.AutoMigrate(
		&Startup{},
		&Tournament{},
		&StartupsTournaments{},
		&Battle{},
		&Events{},
		&BattlesEvents{},
	)
}
