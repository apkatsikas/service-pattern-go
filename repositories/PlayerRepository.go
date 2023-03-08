package repositories

import (
	"github.com/irahardianto/service-pattern-go/customerrors"
	"github.com/irahardianto/service-pattern-go/interfaces"
	"github.com/irahardianto/service-pattern-go/models"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type PlayerRepository struct {
	interfaces.IDbHandler
}

func (repository *PlayerRepository) GetPlayerByName(name string) (*models.PlayerModel, error) {
	gormConn := repository.Connection()

	var player models.PlayerModel
	result := gormConn.First(&player, "name = ?", name)

	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, customerrors.RecordNotFoundError
		}
		return &player, result.Error
	}
	return &player, nil
}
