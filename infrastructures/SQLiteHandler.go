package infrastructures

import (
	"database/sql"

	"gorm.io/gorm"
)

type SQLiteHandler struct {
	// TODO - package private?
	Conn *gorm.DB
}

func (handler *SQLiteHandler) Connection() *gorm.DB {
	return handler.Conn
}

type SqliteRow struct {
	Rows *sql.Rows
}

func (r SqliteRow) Scan(dest ...interface{}) error {
	err := r.Rows.Scan(dest...)
	if err != nil {
		return err
	}

	return nil
}

func (r SqliteRow) Next() bool {
	return r.Rows.Next()
}
