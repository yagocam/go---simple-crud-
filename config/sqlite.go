package config

import (
	"awesomeProject/schemas"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Infof("Creating database at %s", dbPath)
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create("./db/main.db")
		if err != nil {
			return nil, err
		}
		err = file.Close()
		if err != nil {
			return nil, err
		}
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("error opening sqlite db: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&schemas.Post{})
	if err != nil {
		logger.Errorf("error auto-migrating db: %v", err)
		return nil, err
	}
	return db, nil
}
