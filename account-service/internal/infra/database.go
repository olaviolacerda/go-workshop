package infra

import (
	"github.com/olaviolacerda/account/internal/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	conn *gorm.DB
)

func InitDB() error {
	var err error

	dbFile := "./" + GetConfig("db_file")
	conn, err = gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		return err
	}

	err = conn.AutoMigrate(&domain.Account{})
	if err != nil {
		return err
	}

	return nil
}

func Connection() *gorm.DB {
	return conn
}
