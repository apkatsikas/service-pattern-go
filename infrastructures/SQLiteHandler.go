package infrastructures

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLiteHandler struct {
	conn *gorm.DB
}

func (handler *SQLiteHandler) Connection() *gorm.DB {
	return handler.conn
}

func (handler *SQLiteHandler) ConnectSQLite(dsn string) error {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	handler.conn = db
	return nil
}
