package repositories

import (
	ce "github.com/irahardianto/service-pattern-go/customerrors"
	"github.com/irahardianto/service-pattern-go/interfaces"
	"github.com/irahardianto/service-pattern-go/models"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type PlayerRepository struct {
	interfaces.IDbHandler
}

func (pr *PlayerRepository) GetPlayerByName(name string) (*models.Player, error) {
	gormConn := pr.Connection()

	var player models.Player
	result := gormConn.First(&player, "name = ?", name)

	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return &player, ce.ErrRecordNotFound
		}
		return nil, result.Error
	}
	return &player, nil
}
