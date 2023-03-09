package main

import (
	"sync"

	"github.com/irahardianto/service-pattern-go/controllers"
	"github.com/irahardianto/service-pattern-go/infrastructures"
	"github.com/irahardianto/service-pattern-go/repositories"
	"github.com/irahardianto/service-pattern-go/services"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type IServiceContainer interface {
	InjectPlayerController() controllers.PlayerController
}

type kernel struct{}

func (k *kernel) InjectPlayerController() controllers.PlayerController {

	sqliteHandler := &infrastructures.SQLiteHandler{}

	// TODO - movie this into sql lite handler?
	db, _ := gorm.Open(sqlite.Open("/var/tmp/tennis.db"), &gorm.Config{})
	sqliteHandler.Conn = db

	playerService := &services.PlayerService{IPlayerRepository: &repositories.PlayerRepository{IDbHandler: sqliteHandler}}
	playerController := controllers.PlayerController{IPlayerService: playerService}

	return playerController
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
