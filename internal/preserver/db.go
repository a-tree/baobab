package preserver

import (
	"baobab/internal/config"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB(cfg *config.Config) (*gorm.DB, error) {
	var dialector gorm.Dialector
	env := os.Getenv("APP_ENV")

	if env == "production" {
		// Postgres
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.Port)
		dialector = postgres.Open(dsn)
	} else {
		// Development SQLite
		dialector = sqlite.Open("dev.db")
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err == nil {
		db.AutoMigrate(&UserDB{}, &PlaceDB{}) // テーブル自動作成
	}
	return db, err
}
