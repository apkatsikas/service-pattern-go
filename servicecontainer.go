package main

import (
	"fmt"
	"sync"

	"github.com/irahardianto/service-pattern-go/controllers"
	"github.com/irahardianto/service-pattern-go/infrastructures"
	"github.com/irahardianto/service-pattern-go/repositories"
	"github.com/irahardianto/service-pattern-go/services"
)

const dbPath = "/var/tmp/tennis.db"

type IServiceContainer interface {
	InjectPlayerController() controllers.PlayerController
}

type kernel struct{}

func (k *kernel) InjectPlayerController() controllers.PlayerController {

	sqliteHandler := &infrastructures.SQLiteHandler{}

	err := sqliteHandler.ConnectSQLite(dbPath)
	if err != nil {
		panic(err)
	}

	if migrateDB {
		fmt.Println("Migrating DB...")
		err = sqliteHandler.Migrate()
		if err != nil {
			errMsg := fmt.Errorf("Got an error on migration %v", err)
			fmt.Println(errMsg.Error())
		}
	}

	playerService := &services.PlayerService{IPlayerRepository: &repositories.PlayerRepository{IDbHandler: sqliteHandler}}
	playerController := controllers.PlayerController{IPlayerService: playerService}

	return playerController
}

var (
	k             *kernel
	containerOnce sync.Once
)

// singleton
func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
