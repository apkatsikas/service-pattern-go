package infrastructures

import (
	"github.com/irahardianto/service-pattern-go/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLiteHandler struct {
	conn *gorm.DB
}

func (handler *SQLiteHandler) migrateCreate(name string, score int) *gorm.DB {
	var player models.PlayerModel
	return handler.conn.Where(
		models.PlayerModel{Name: name}).Attrs(
		models.PlayerModel{Score: score}).FirstOrCreate(&player)
}

func (handler *SQLiteHandler) migrate() error {
	// Auto Migrate
	err := handler.conn.AutoMigrate(&models.PlayerModel{})
	if err != nil {
		return err
	}

	// Setup data
	result := handler.migrateCreate("Rafael", 3)

	if result.Error != nil {
		return result.Error
	}

	result = handler.migrateCreate("Roger", 2)

	if result.Error != nil {
		return result.Error
	}

	result = handler.migrateCreate("Serena", 1)

	if result.Error != nil {
		return result.Error
	}

	result = handler.migrateCreate("Maria", 0)

	if result.Error != nil {
		return result.Error
	}

	return nil
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
	err = handler.migrate()
	if err != nil {
		return err
	}
	return nil
}
