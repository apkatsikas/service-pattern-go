package infrastructures

import (
	"github.com/irahardianto/service-pattern-go/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLiteHandler struct {
	conn *gorm.DB
}

func (handler *SQLiteHandler) migrateCreate(player models.Player) error {
	var p models.Player
	// If record can't be found, insert it
	result := handler.conn.Where(
		models.Player{Name: player.Name}).Attrs(
		models.Player{Score: player.Score}).FirstOrCreate(&p)

	return result.Error
}

func getMigrationData() []models.Player {
	players := []models.Player{
		{
			Name:  "Rafael",
			Score: 3,
		},
		{
			Name:  "Roger",
			Score: 2,
		},
		{
			Name:  "Serena",
			Score: 1,
		},
		{
			Name:  "Maria",
			Score: 0,
		},
	}
	return players
}

func (handler *SQLiteHandler) Connection() *gorm.DB {
	return handler.conn
}

func (handler *SQLiteHandler) Migrate() error {
	// Create table if needed
	err := handler.conn.AutoMigrate(&models.Player{})
	if err != nil {
		return err
	}

	// Loop through and insert migration data
	for _, p := range getMigrationData() {
		err = handler.migrateCreate(p)
		if err != nil {
			return err
		}
	}

	return nil
}

func (handler *SQLiteHandler) ConnectSQLite(dsn string) error {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		// TODO - turn this on?
		//Logger: logger.Default.LogMode(logger.Silent)
	})
	if err != nil {
		return err
	}
	handler.conn = db
	return nil
}
