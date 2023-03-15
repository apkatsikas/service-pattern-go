package main

import (
	"sync"

	"github.com/irahardianto/service-pattern-go/controllers"
	"github.com/irahardianto/service-pattern-go/infrastructures"
	"github.com/irahardianto/service-pattern-go/infrastructures/logutil"
	"github.com/irahardianto/service-pattern-go/repositories"
	"github.com/irahardianto/service-pattern-go/services"
)

const dbPath = "/var/tmp/tennis.db"

type IServiceContainer interface {
	InjectPlayerController() controllers.PlayerController
}

type kernel struct{}

func (k *kernel) InjectPlayerController() controllers.PlayerController {
	// Setup sqlite
	sqliteHandler := &infrastructures.SQLiteHandler{}

	// Connect to SQLite
	err := sqliteHandler.ConnectSQLite(dbPath)
	if err != nil {
		logutil.Fatal("Failed to connect to SQLite. Error was %v", err)
	}

	// Migrate DB if flag set
	if FlagUtil().Get().MigrateDB {
		logutil.Info("Migrating DB...")
		err = sqliteHandler.Migrate()
		if err != nil {
			logutil.Error("Failed to migrate. Error was %v", err)
		}
	}

	// Inject Player service with repository
	playerService := &services.PlayerService{
		IPlayerRepository: &repositories.PlayerRepository{IDbHandler: sqliteHandler}}
	// Inject controller with Player service
	playerController := controllers.PlayerController{IPlayerService: playerService}

	return playerController
}

// Setup singleton
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
