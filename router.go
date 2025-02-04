package main

import (
	"sync"

	"github.com/go-chi/chi"
	"github.com/irahardianto/service-pattern-go/infrastructures/logutil"
)

type IChiRouter interface {
	InitRouter() *chi.Mux
}

type router struct{}

func (router *router) InitRouter() *chi.Mux {
	// Setup controller
	playerController := ServiceContainer().InjectPlayerController()

	// Create router
	r := chi.NewRouter()
	r.HandleFunc("/getScore/{player1}/vs/{player2}", playerController.GetPlayerScore)

	logutil.Info("Router initialized.")

	return r
}

// Setup singleton
var (
	m          *router
	routerOnce sync.Once
)

func ChiRouter() IChiRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}
	return m
}
